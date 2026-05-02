package optimize

import "os/exec"

type exploerBackground struct{}

func (e *exploerBackground) Do() error {
	res := exec.Command("reg", "add", "\"HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\"", "/v", "TaskbarAcrylicOpacity", "/t", "REG_DWORD", "/d", "0xc8", "/f")
	err := res.Run()
	return err
}

func (e *exploerBackground) Cancel() error {
	res := exec.Command("reg", "add", "\"HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\"", "/v", "TaskbarAcrylicOpacity", "/t", "REG_DWORD", "/d", "0xc8", "/f")
	err := res.Run()
	return err
}
