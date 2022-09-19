package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

type Inis struct {
	ServerProtocol string
	ServerIp       string
	ServerPort     string
}

func CheckIni(filePath string, inis *Inis) {

	cgf, err := ini.Load(filePath)
	if err != nil {
		fmt.Println("[!] setting.ini File not exist!")
		time.Sleep(10 * time.Minute)
	}
	inis.ServerIp = cgf.Section("server").Key("ip").String()
	inis.ServerPort = cgf.Section("server").Key("port").String()
	inis.ServerProtocol = cgf.Section("server").Key("protocol").String()

	if inis.ServerIp == "" || inis.ServerPort == "" || inis.ServerProtocol == "" {
		if inis.ServerIp == "" {
			fmt.Println("[!] Check serverIp setting.ini File")
		}
		if inis.ServerPort == "" {
			fmt.Println("[!] Check serverPort setting.ini File")
		}

		if inis.ServerProtocol == "" {
			fmt.Println("[!] Check serverProtocol setting.ini File")
		}
		fmt.Println("[!] Read File Fail!")
		fmt.Println("Exit AG-Installer 10s...!")
		time.Sleep(10 * time.Second)
		os.Exit(3)

	}

}
