package monitor

import (
	"api/manager/events"
	"bytes"
	"fmt"
	"os/exec"
)

type MANAGER struct {
}

//var event events.Event
// manager进程监控
func (m *MANAGER) ProcessMonitor() (string, error) {
	var output bytes.Buffer
	cmd := exec.Command("ps", "aux ", " | grep lotus", "| grep -v grep", "| wc -l")
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		fmt.Println("执行获取manager进程失败：", err)
	}
	switch output.String() {
	case "0":
		//通知事件侦听器
		events.Trigger(events.NewEvent("manager process abnormal", 0))
		//return output.String(),error()
	case "1":
		return output.String(), nil
	}
	return "", nil
}
