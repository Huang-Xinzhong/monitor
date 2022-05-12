package monitor

import (
	"api/manager/baseInfo"
	"fmt"
	"os"
	_ "os"
)

type PROFILE struct {
}

//const (
//	lotusrepo = "/mnt/lotus/"
//	lotusbin = "/mnt/lotus/lotus/"
//	TMPDIR = "/mnt/lotus/tmp/"
//)

var operatingSystem = baseInfo.GetOS()
var user = baseInfo.GetUserName()
var Home = func() string {
	var home string
	if operatingSystem == "Ubuntu" {
		home = fmt.Sprintf("/home/%s/", operatingSystem)
		fmt.Println(home)
	} else if operatingSystem == "Centos" {
		home = "/root/"
	}
	return home
}

func getProfile() {
	lotusrepo := os.Getenv("lotusrepo")
	lotusbin := os.Getenv("lotusbin")
	TMPDIR := os.Getenv("TMPDIR")

	fmt.Println(lotusrepo)
	fmt.Println(lotusbin)
	fmt.Println(TMPDIR)

}
