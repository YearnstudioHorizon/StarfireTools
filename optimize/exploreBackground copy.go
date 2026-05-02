package optimize

import (
	"fmt"
	"os"
)

type exploerBackground1 struct{}

func (e *exploerBackground1) Do() error {
	fmt.Println("生效")
	// res := exec.Command("cmd", "/c", "\"echo 0000 > %TEMP%/StarfireTool-elevate.log\"")
	// err := res.Run()
	_, err := os.Create("C:/Users/33032/Documents/编程/StarFireTool/build/bin/temp.log")
	return err
}

func (e *exploerBackground1) Cancel() error {
	fmt.Println("取消生效")
	// res := exec.Command("cmd", "/c", "\"echo 1111 > %TEMP%/StarfireTool-elevate.log\"")
	// err := res.Run()
	_, err := os.Create("C:/Users/33032/Documents/编程/StarFireTool/build/bin/temp1.log")
	return err
}

func (e *exploerBackground1) GetStatus() (bool, error) {
	return false, nil
}
func (e *exploerBackground1) GetName() string {
	return "测试优化项"
}

func (e *exploerBackground1) GetDesc() string {
	return "测试说明"
}

func NewExploreBackground1() *exploerBackground1 {
	return &exploerBackground1{}
}
