package events

import (
	"fmt"
)

// manager进程异常处理
func ManagerProcessHandler(event Event) {
	fmt.Println("manager process abnormal, process number is ", 0)
}
