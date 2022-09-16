package main

import (
	"AG-Installer/utils"
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	var osv string
	cgf, err := ini.Load("setting.ini")
	if err != nil {
		fmt.Println("[!] setting.ini File not exist!")
		time.Sleep(10 * time.Minute)
	}

	v := runtime.GOOS
	switch v {
	case "windows":
		osv = "windows"
	case "linux":
		osv = "linux"
	}
	fmt.Println("[-] OS Type", osv)

	serverIp := cgf.Section("server").Key("ip").String()
	serverPort := cgf.Section("server").Key("port").String()
	serverProtocol := cgf.Section("server").Key("protocol").String()
	if serverIp == "" || serverPort == "" || serverProtocol == "" {
		if serverIp == "" {
			fmt.Println("[!] Check serverIp setting.ini File")
		}
		if serverPort == "" {
			fmt.Println("[!] Check serverPort setting.ini File")
		}

		if serverProtocol == "" {
			fmt.Println("[!] Check serverProtocol setting.ini File")
		}
		fmt.Println("[!] Read File Fail!")
		fmt.Println("Exit AG-Installer 10s...!")
		time.Sleep(10 * time.Second)
		os.Exit(3)

	}
	fmt.Println("[-] Server IP:PORT", serverIp, ":", serverPort)

	checkURL := serverProtocol + "://" + serverIp + ":" + serverPort + "/" + osv + "/" + "AG-Agent-List"
	fmt.Println("[-] Connecting to Server...(10s)")
	fmt.Println("[-] Waiting For", checkURL)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(checkURL)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("[!] Agent List Download Fail!")
		fmt.Println("Exit AG-Installer 10s...!")
		time.Sleep(10 * time.Second)
		os.Exit(3)
	} else {
		fmt.Println("[O] Agent List Get Success!")
	}

	data, _ := io.ReadAll(resp.Body)
	agentInfos := strings.Split(string(data), "\r\n")
	agentInfo := strings.Split(agentInfos[0], " ")
	aName := agentInfo[0]
	aVer := agentInfo[1]

	downURL := serverProtocol + "://" + serverIp + ":" + serverPort + "/" + osv + "/AG-Agent/" + aVer + "/" + aName
	fmt.Println("[-] Connecting to Server...(10s)")
	fmt.Println("[-] Downloading...", aName, aVer)
	value := os.Getenv("ProgramFiles")

	err = os.Mkdir(value+"\\Anti-Gravity", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	path, _ := utils.Downloads(downURL, value+"\\"+aName)
	fmt.Println(path)

}
