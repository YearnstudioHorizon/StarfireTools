## Starfire Tools

> 一个聚合系统工具与功能管理的桌面应用

### 技术栈

- Wails
- Vue3
- Vite

### 快速启动

```bash
wails dev
```

### 许可证

Apache 2.0

### CI/CD

流水线自动构建并发布产物, 请使用语义化commit msg, 如 `feat: ...、fix: ...`

### 添加优化项

#### 添加优化项

1. 在`optimize`文件夹下新建一个`.go`文件, 声明为`optimize`包
2. 实现位于`optimize/interface.go`的`OptimizeItem`接口, 接口的注释表明了每个函数的作用
3. 提供一个构造函数, 返回实现接口的结构体的指针
4. 将这个优化项添加到目标组内(请看下一个条目"添加优化组")

示例代码
```go
package optimize

import (
	"fmt"
)

type test1 struct{} // 空结构体即可

func (e *test1) Do() error {
    // 此函数会在启用该功能时被调用
	fmt.Println("生效")
	return err
}

func (e *test1) Cancel() error {
    // 此函数会在关闭该功能时被调用
	fmt.Println("取消生效")
	return err
}

func (e *test1) GetStatus() (bool, error) {
    // 返回当前状态
	return false, nil
}
func (e *test1) GetName() string {
    // 返回显示的名称
	return "优化项名称"
}

func (e *test1) GetDesc() string {
    // 返回显示的说明
	return "优化项说明"
}

// 构造函数
func NewTest1() *test1 {
	return &test1{} // 必须是指针
}

```

#### 添加优化组

在`app.go`内找到

```go
var ItemGroups []ItemGroup
```

然后添加一行函数调用

```go
ItemGroups = AppendNewGroup(ItemGroups, "优化组名", []optimize.OptimizeItem{optimize.构造函数1(), optimize.构造函数2()})
```

这里的`构造函数1`即为优化组件对应的构造函数名

如果要将一个已有的加入现有优化组, 在`[]optimize.OptimizeItem`中添加对应的构造函数即可

#### 参考接口定义 (optimize/interface.go)

> ```go
> type OptimizeItem interface {
>	Do() error                // 在功能被打开时被调用
>	Cancel() error            // 在功能被关闭时被调用
>	GetStatus() (bool, error) // 获取当前开启关闭状态
>	GetName() string          // 获取名称
>	GetDesc() string          // 获取描述
> }
> ```