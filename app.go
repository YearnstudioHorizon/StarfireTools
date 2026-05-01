package main

import (
	"context"
	"os/user"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx           context.Context
	Username      string
	OS            string
	KernelVersion string
	CPUName       string
	CPUMHz        float64
	Cores         int
	PhyMem        uint64
	VirMem        uint64
	Version       string
}

// NewApp creates a new App application struct
func NewApp() *App {
	u, _ := user.Current()
	sys, _ := host.Info()
	cu, _ := cpu.Info()
	memory, _ := mem.VirtualMemory()
	vmem, _ := mem.SwapMemory()
	phyMem := memory.Total
	vMem := vmem.Total
	var CPUName string = cu[0].ModelName
	var CPUMHz float64 = cu[0].Mhz
	CPUCore := cu[0].Cores
	return &App{
		Username:      u.Username,
		OS:            sys.Platform,
		KernelVersion: sys.KernelVersion,
		CPUName:       CPUName,
		CPUMHz:        CPUMHz,
		Cores:         int(CPUCore),
		PhyMem:        phyMem,
		VirMem:        vMem,
		Version:       VERSION,
	}
}

const VERSION string = "0.1.0 α"

func (a *App) GetUserName() string {
	return a.Username
}

func (a *App) GetOSName() string {
	return a.OS
}

func (a *App) GetKernelVersion() string {
	return a.KernelVersion
}

func (a *App) GetCPU() string {
	return a.CPUName
}

func (a *App) GetCPUCores() int {
	return a.Cores
}

func (a *App) GetCPUMHz() float64 {
	return a.CPUMHz
}

func (a *App) GetPhyMem() uint64 {
	return a.PhyMem
}

func (a *App) GetVirMem() uint64 {
	return a.VirMem
}

func (a *App) GetVersion() string {
	return a.Version
}

func (a *App) startup(ctx context.Context) {
	_, err := user.Current()
	if err != nil {
		runtime.Quit(ctx)
	}
	a.ctx = ctx
	runtime.WindowCenter(ctx)
	// runtime.EventsOn(ctx, "wails:window-maximise", func(optionalData ...interface{}) {
	// 	runtime.LogInfo(ctx, "窗口最大化")
	// })
}
