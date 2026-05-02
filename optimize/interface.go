package optimize

type OptimizeItem interface {
	Do() error                // 在功能被打开时被调用
	Cancel() error            // 在功能被关闭时被调用
	GetStatus() (bool, error) // 获取当前开启关闭状态
	GetName() string          // 获取名称
	GetDesc() string          // 获取描述
}
