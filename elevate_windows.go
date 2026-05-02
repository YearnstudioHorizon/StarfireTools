//go:build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	elevateHelperArg = "--elevate-to-system-helper"

	// Token 权限
	StandardRightsRequired = 0x000F0000
	TokenAssignPrimary     = 0x0001
	TokenDuplicate         = 0x0002
	TokenImpersonate       = 0x0004
	TokenQuery             = 0x0008
	TokenQuerySource       = 0x0010
	TokenAdjustPrivileges  = 0x0020
	TokenAdjustGroups      = 0x0040
	TokenAdjustDefault     = 0x0080
	TokenAdjustSessionId   = 0x0100
	TokenAllAccess         = StandardRightsRequired | TokenAssignPrimary | TokenDuplicate | TokenImpersonate | TokenQuery | TokenQuerySource | TokenAdjustPrivileges | TokenAdjustGroups | TokenAdjustDefault | TokenAdjustSessionId

	// 进程权限
	ProcessQueryInformation = 0x0400
)

var (
	modadvapi32                 = syscall.NewLazyDLL("advapi32.dll")
	procOpenProcessToken        = modadvapi32.NewProc("OpenProcessToken")
	procDuplicateTokenEx        = modadvapi32.NewProc("DuplicateTokenEx")
	procCreateProcessWithTokenW = modadvapi32.NewProc("CreateProcessWithTokenW")
)

func logElevation(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logPath := filepath.Join(os.TempDir(), "StarfireTool-elevate.log")
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		timestamp := time.Now().Format("2006-01-02 15:04:05.000")
		f.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, msg))
	}
}

func maybeRelaunchAsSystem() (bool, error) {
	logElevation("=== maybeRelaunchAsSystem started ===")
	if shouldSkipAutoElevation() {
		logElevation("Skipping auto elevation")
		return false, nil
	}

	isSystem, err := runningAsSystem()
	if err == nil && isSystem {
		logElevation("Already running as SYSTEM")
		return false, nil
	}

	exePath, err := os.Executable()
	if err != nil {
		logElevation("Failed to get executable path: %v", err)
		return false, err
	}
	logElevation("Exe path: %s", exePath)
	// Avoid trying to elevate when running via `go run` / dev builds (temp exe path).
	// In those cases the executable path can be ephemeral and task launch may fail.
	if looksLikeGoRunTempExe(exePath) {
		logElevation("Looks like go run temp exe, skipping")
		return false, nil
	}

	helper, passThroughArgs := parseElevationArgs(os.Args[1:])
	logElevation("Parsed args: helper=%v, passThrough=%v", helper, passThroughArgs)

	if helper {
		logElevation("Running as helper, attempting token elevation")
		err = runAsSystemViaToken(exePath, passThroughArgs)
		if err != nil {
			logElevation("Helper failed to launch SYSTEM process: %v", err)
			return true, fmt.Errorf("helper failed to launch SYSTEM process: %v", err)
		}
		logElevation("Helper successfully launched SYSTEM process")
		return true, nil
	}

	// 尝试直接 Token 提权（当前用户如果已经是管理员，会直接成功）
	logElevation("Attempting direct token elevation")
	err = runAsSystemViaToken(exePath, passThroughArgs)
	if err == nil {
		logElevation("Direct token elevation succeeded")
		return true, nil
	}
	logElevation("Direct token elevation failed: %v", err)

	// 若上面失败，代表目前只是普通用户，需要拉起带有 Helper 参数的管理员进程触发 UAC
	logElevation("Starting self elevated via PowerShell")
	args := append([]string{elevateHelperArg}, passThroughArgs...)

	err = startSelfElevated(exePath, args)
	if err != nil {
		logElevation("startSelfElevated failed: %v", err)
		return false, err
	}
	logElevation("startSelfElevated succeeded (UAC triggered)")
	return true, nil
}

func shouldSkipAutoElevation() bool {
	if v := strings.TrimSpace(os.Getenv("STARFIRE_NO_SYSTEM")); v != "" && v != "0" && strings.ToLower(v) != "false" {
		return true
	}
	// Best-effort skip in `wails dev` / hot-reload scenarios.
	devEnvKeys := []string{
		"devserver",
		"frontenddevserverurl",
		"assetdir",
		"WAILS_DEV",
		"WAILS_DEBUG",
		"WAILS_LIVERELOAD",
	}
	for _, k := range devEnvKeys {
		if strings.TrimSpace(os.Getenv(k)) != "" {
			return true
		}
	}
	for _, a := range os.Args[1:] {
		la := strings.ToLower(a)
		if la == "-dev" || la == "--dev" || strings.Contains(la, "wails") && strings.Contains(la, "dev") {
			return true
		}
	}
	return false
}

func looksLikeGoRunTempExe(exePath string) bool {
	p := strings.ToLower(exePath)
	return strings.Contains(p, `\appdata\local\temp\go-build`) ||
		strings.Contains(p, `\appdata\local\go-build`) ||
		strings.Contains(p, `\tmp\go-build`) ||
		strings.Contains(p, `wailsbindings`)
}

func parseElevationArgs(args []string) (helper bool, passThrough []string) {
	for _, arg := range args {
		if arg == elevateHelperArg {
			helper = true
		} else {
			passThrough = append(passThrough, arg)
		}
	}
	return helper, passThrough
}

func startSelfElevated(exePath string, args []string) error {
	psArgs := []string{
		"-WindowStyle", "Hidden", // 隐藏 PowerShell 自身的闪烁窗口
		"-NoProfile",
		"-NonInteractive",
		"-Command",
		buildStartProcessCommand(exePath, args),
	}
	cmd := exec.Command("powershell", psArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func buildStartProcessCommand(exePath string, args []string) string {
	quotedArgs := make([]string, 0, len(args))
	for _, a := range args {
		quotedArgs = append(quotedArgs, fmt.Sprintf("'%s'", escapePowerShellSingleQuoted(a)))
	}
	return fmt.Sprintf("Start-Process -FilePath '%s' -ArgumentList %s -Verb RunAs",
		escapePowerShellSingleQuoted(exePath),
		strings.Join(quotedArgs, ", "),
	)
}

func escapePowerShellSingleQuoted(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func runAsSystemViaToken(exePath string, args []string) error {
	logElevation("runAsSystemViaToken started for %s", exePath)
	var token windows.Token
	err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token)
	if err != nil {
		logElevation("OpenProcessToken failed: %v", err)
		return fmt.Errorf("OpenProcessToken failed: %w", err)
	}
	defer token.Close()

	err = enablePrivilege(token, "SeDebugPrivilege")
	errImp := enablePrivilege(token, "SeImpersonatePrivilege") // 为 CreateProcessWithTokenW 提供保证
	logElevation("Privileges adjusted (SeDebugPrivilege: %v, SeImpersonatePrivilege: %v)", err == nil, errImp == nil)

	hDuplicatedToken, err := getSystemToken()
	if err != nil {
		logElevation("getSystemToken failed: %v", err)
		return err
	}
	defer hDuplicatedToken.Close()

	// 将拷贝来的 Token 设置到当前活动的 Console 会话 ID 中，以保证界面能够正常在当前用户桌面上绘制
	sessionID := windows.WTSGetActiveConsoleSessionId()
	var currentTokenSessionId uint32
	var returnLen uint32
	err = windows.GetTokenInformation(hDuplicatedToken, 12, (*byte)(unsafe.Pointer(&currentTokenSessionId)), 4, &returnLen)

	// 仅当实在找不到相同 Session 的 SYSTEM token (使用了备选 token)，才做 SetTokenInformation
	if err == nil && currentTokenSessionId != sessionID {
		logElevation("Token Session ID (%d) != Active Console Session ID (%d). Attempting to fix...", currentTokenSessionId, sessionID)
		err = windows.SetTokenInformation(
			hDuplicatedToken,
			12, // windows.TokenSessionId
			(*byte)(unsafe.Pointer(&sessionID)),
			uint32(unsafe.Sizeof(sessionID)),
		)
		if err != nil {
			logElevation("SetTokenInformation failed (ignoring): %v", err)
		} else {
			logElevation("Token Session ID successfully changed")
		}
	}

	var si windows.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	si.Desktop = windows.StringToUTF16Ptr("winsta0\\default")

	var pi windows.ProcessInformation
	cmdLine := makeCmdLine(exePath, args)
	logElevation("Command line: %s", cmdLine)
	commandLinePtr := windows.StringToUTF16Ptr(cmdLine)

	r, _, err := procCreateProcessWithTokenW.Call(
		uintptr(hDuplicatedToken),
		1, // LOGON_WITH_PROFILE (0x00000001) - 允许系统加载目标用户的配置文件
		0,
		uintptr(unsafe.Pointer(commandLinePtr)),
		0,
		0,
		0,
		uintptr(unsafe.Pointer(&si)),
		uintptr(unsafe.Pointer(&pi)),
	)
	if r == 0 {
		logElevation("CreateProcessWithTokenW failed: %v", err)
		return fmt.Errorf("CreateProcessWithTokenW failed: %v", err)
	}

	logElevation("Successfully created process with SYSTEM privilege, new PID: %d", pi.ProcessId)
	windows.CloseHandle(pi.Process)
	windows.CloseHandle(pi.Thread)
	return nil
}

func makeCmdLine(exePath string, args []string) string {
	var sb strings.Builder
	sb.WriteString(syscall.EscapeArg(exePath))
	for _, arg := range args {
		sb.WriteString(" ")
		sb.WriteString(syscall.EscapeArg(arg))
	}
	return sb.String()
}

// getSystemToken 自动遍历进程并复制最适合当前会话的 SYSTEM 进程 Token
func getSystemToken() (windows.Token, error) {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return 0, err
	}
	defer windows.CloseHandle(snapshot)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	targetSessionId := windows.WTSGetActiveConsoleSessionId()
	systemSid, err := windows.StringToSid("S-1-5-18")
	if err != nil {
		return 0, err
	}

	var fallbackToken windows.Token
	var fallbackPid uint32

	err = windows.Process32First(snapshot, &entry)
	for err == nil {
		pid := entry.ProcessID
		err = windows.Process32Next(snapshot, &entry)
		if pid == 0 {
			continue
		}

		// 使用较低的请求权限：PROCESS_QUERY_LIMITED_INFORMATION (0x1000) 彻底绕过受保护进程的 Access Denied 拦截
		hProcess, errProcess := windows.OpenProcess(0x1000, false, pid)
		if errProcess != nil {
			continue
		}

		var hToken windows.Token
		errToken := windows.OpenProcessToken(hProcess, windows.TOKEN_QUERY|windows.TOKEN_DUPLICATE, &hToken)
		windows.CloseHandle(hProcess)
		if errToken != nil {
			continue
		}

		user, errUser := hToken.GetTokenUser()
		if errUser != nil || !user.User.Sid.Equals(systemSid) {
			hToken.Close()
			continue
		}

		var sessionID uint32
		var returnLen uint32
		errSession := windows.GetTokenInformation(hToken, 12, (*byte)(unsafe.Pointer(&sessionID)), 4, &returnLen)

		// 如果找到了完全处于当前界面的 SYSTEM (完美符合，避免后面需要 SeTcbPrivilege 强改 Session)
		if errSession == nil && sessionID == targetSessionId {
			var hDuplicatedToken windows.Token
			r, _, _ := procDuplicateTokenEx.Call(
				uintptr(hToken),
				uintptr(TokenAllAccess),
				0, 3, 1,
				uintptr(unsafe.Pointer(&hDuplicatedToken)),
			)
			hToken.Close()
			if r != 0 {
				logElevation("Found perfect SYSTEM token from PID: %d (Session %d)", pid, sessionID)
				if fallbackToken != 0 {
					fallbackToken.Close()
				}
				return hDuplicatedToken, nil
			}
		} else {
			// 备选项：如果是 SYSTEM 但在别的 Session (比如 Session 0), 记录下来作为备选
			if fallbackToken == 0 {
				r, _, _ := procDuplicateTokenEx.Call(uintptr(hToken), uintptr(TokenAllAccess), 0, 3, 1, uintptr(unsafe.Pointer(&fallbackToken)))
				if r != 0 {
					fallbackPid = pid
				}
			}
			hToken.Close()
		}
	}

	if fallbackToken != 0 {
		logElevation("Using fallback SYSTEM token from PID: %d, will attempt to set SessionID", fallbackPid)
		return fallbackToken, nil
	}

	return 0, fmt.Errorf("could not find any SYSTEM token to duplicate")
}

func enablePrivilege(token windows.Token, privilege string) error {
	var luid windows.LUID
	err := windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr(privilege), &luid)
	if err != nil {
		return err
	}

	newState := windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges: [1]windows.LUIDAndAttributes{
			{
				Luid:       luid,
				Attributes: windows.SE_PRIVILEGE_ENABLED,
			},
		},
	}

	var returnLength uint32
	return windows.AdjustTokenPrivileges(
		token,
		false,
		&newState,
		uint32(unsafe.Sizeof(newState)),
		nil,
		&returnLength,
	)
}

func runningAsSystem() (bool, error) {
	var token windows.Token
	err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token)
	if err != nil {
		return false, err
	}
	defer token.Close()

	user, err := token.GetTokenUser()
	if err != nil {
		return false, err
	}
	systemSid, err := windows.StringToSid("S-1-5-18")
	if err != nil {
		return false, err
	}
	return user.User.Sid.Equals(systemSid), nil
}
