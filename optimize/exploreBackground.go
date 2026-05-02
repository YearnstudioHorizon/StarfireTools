package optimize

import (
	"golang.org/x/sys/windows/registry"
)

type exploerBackground struct{}

func (e *exploerBackground) Do() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	// 0xc8 对应的十进制是 200
	return k.SetDWordValue("TaskbarAcrylicOpacity", 0xc8)
}

func (e *exploerBackground) Cancel() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.DeleteValue("TaskbarAcrylicOpacity")
	// 如果注册表值本来就不存在，我们将其视为已成功取消，不抛出错误
	if err != nil && err != registry.ErrNotExist {
		return err
	}
	return nil
}

func (e *exploerBackground) GetStatus() (bool, error) {
	// fmt.Print("已执行查询命令")
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.QUERY_VALUE)
	if err != nil {
		return false, nil
	}
	defer k.Close()

	val, _, err := k.GetIntegerValue("TaskbarAcrylicOpacity")
	if err != nil {
		return false, nil
	}

	// 检查读出的值是否匹配 0xc8 (200)
	return val == 0xc8, nil
}
func (e *exploerBackground) GetName() string {
	return "任务栏透明"
}

func (e *exploerBackground) GetDesc() string {
	return "仅Windows10可用"
}

func NewExploreBackground() *exploerBackground {
	return &exploerBackground{}
}
