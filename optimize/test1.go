package optimize

import (
	"fmt"
	"os"
)

type test1 struct{}

func (e *test1) Do() error {
	fmt.Println("生效")
	_, err := os.Create("C:/Users/33032/Documents/编程/StarFireTool/build/bin/temp.log")
	return err
}

func (e *test1) Cancel() error {
	fmt.Println("取消生效")
	_, err := os.Create("C:/Users/33032/Documents/编程/StarFireTool/build/bin/temp1.log")
	return err
}

func (e *test1) GetStatus() (bool, error) {
	return false, nil
}
func (e *test1) GetName() string {
	return "测试优化项"
}

func (e *test1) GetDesc() string {
	return "测试说明"
}

func NewTest1() *test1 {
	return &test1{}
}
