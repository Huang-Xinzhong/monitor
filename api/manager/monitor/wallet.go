package monitor

import (
	"bytes"
	"fmt"
	"os/exec"
)

func WalletMonitor() {
	var output bytes.Buffer
	cmd := exec.Command()
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		fmt.Println("执行获取manager进程失败：", err)
	}
}
