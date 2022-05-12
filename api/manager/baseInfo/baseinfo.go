package baseInfo

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
)

type Info struct {
	ipAddr    string //本机地址
	os        string //操作系统
	osVersion string //操作系统版本
	userName  string //用户名
}

// 获取内网地址
func (i *Info) obtainIpaddr() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("本机ip地址：", ipnet.IP.String())
				i.ipAddr = ipnet.IP.String()
			}
		}
	}
}
func (i *Info) GetIpaddr() string {
	return i.ipAddr
}

//获取操作系统（centos, ubuntu）
func (i *Info) obtainOS() {
	//var out bytes.Buffer
	out, err := exec.Command("cat", "/etc/issue", "| awk '{print $1}'").Output()
	// Ubuntu 20.04.3 LTS \n \l
	if err != nil {
		fmt.Println(err)
	}
	i.os = string(out)

}
func GetOS() string {
	return info.os
}

// 获取系统版本
func (i *Info) obtainOsVersion() {
	var out bytes.Buffer
	cmd := exec.Command("cat", "/etc/issue", "| awk '{print $2}'")
	// Ubuntu 20.04.3 LTS \n \l
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Printf("获取系统版本失败：%s\n", err)
	}
	i.osVersion = out.String()
}
func GetOsVersion() string {
	return info.osVersion
}

//获取用户名
func (i *Info) obtainUserName() {
	i.userName, _ = os.Hostname()
}
func GetUserName() string {
	return info.userName
}

// 初始化基础信息配置
var info Info

func InfoInit() {
	info.obtainIpaddr()
	info.obtainOS()
	info.obtainOsVersion()
	info.obtainUserName()
}
