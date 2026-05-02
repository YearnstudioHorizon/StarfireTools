package optimize

type OptimizeItem interface {
	Do() error                // 开启
	Cancel() error            // 关闭
	GetStatus() (bool, error) // 获取当前状态
	GetName() string          // 获取名称
	GetDesc() string          // 获取描述
}
