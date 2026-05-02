package main

import (
	"context"
	"fmt"
	"os/user"

	"StarFireTool/optimize"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ItemGroup struct {
	Name      string                  `json:"name"`
	ItemNames []string                `json:"items"`
	ItemDescs []string                `json:"items_descs"`
	Items     []optimize.OptimizeItem `json:"-"`
}

func (group *ItemGroup) addContent(item optimize.OptimizeItem) {
	// fmt.Printf("注册了名为%v的优化项, 描述为%v\n", item.GetName(), item.GetDesc())
	group.ItemNames = append(group.ItemNames, item.GetName())
	group.ItemDescs = append(group.ItemDescs, item.GetDesc())
	group.Items = append(group.Items, item)
}

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
	ItemsGroups   []ItemGroup
}

func AppendNewGroup(target []ItemGroup, name string, items []optimize.OptimizeItem) []ItemGroup {
	group := ItemGroup{Name: name}
	for _, v := range items {
		group.addContent(v)
	}
	return append(target, group)
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
	// 初始化配置组
	var ItemGroups []ItemGroup
	ItemGroups = AppendNewGroup(ItemGroups, "系统效果优化", []optimize.OptimizeItem{optimize.NewExploreBackground(), optimize.NewTest1()})

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
		ItemsGroups:   ItemGroups,
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

func (a *App) GetItemGroups() []ItemGroup {
	return a.ItemsGroups
}

func (a *App) Do(group, item string) error {
	// 查找组
	for _, v := range a.ItemsGroups {
		if v.Name == group {
			// 查到了对应优化组
			for k, value := range v.ItemNames {
				if value == item {
					// 找到了优化项
					return v.Items[k].Do()
				}
			}
			return fmt.Errorf("未找到优化项:%v", group)
		}
	}
	return fmt.Errorf("未找到优化组:%v", group)
}

func (a *App) Cancel(group, item string) error {
	// 查找组
	for _, v := range a.ItemsGroups {
		if v.Name == group {
			// 查到了对应优化组
			for k, value := range v.ItemNames {
				if value == item {
					// 找到了优化项
					return v.Items[k].Cancel()
				}
			}
			return fmt.Errorf("未找到优化项:%v", group)
		}
	}
	return fmt.Errorf("未找到优化组:%v", group)
}

func (a *App) GetStatus(group, item string) (bool, error) {
	// 查找组
	for _, v := range a.ItemsGroups {
		if v.Name == group {
			// 查到了对应优化组
			for k, value := range v.ItemNames {
				if value == item {
					// 找到了优化项
					return v.Items[k].GetStatus()
				}
			}
			return false, fmt.Errorf("未找到优化项:%v", group)
		}
	}
	return false, fmt.Errorf("未找到优化组:%v", group)
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
