package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetAgentListUrl(inis *Inis, osv string, target string) string {
	fmt.Println("[-] OS Type", osv)
	fmt.Println("[-] Server IP:PORT", inis.ServerIp, ":", inis.ServerPort)
	agentList := ""
	switch target {
	case "agent":
		agentList = "AG-Agent-List.txt"
	case "updater":
		agentList = "AG-Updater-List.txt"
	}

	checkURL := inis.ServerProtocol + "://" + inis.ServerIp + ":" + inis.ServerPort + "/" + osv + "/" + agentList

	return checkURL
}

func CheckList(target string, checkURL string) (string, string) {
	fmt.Println("[-] Connecting to Server...(10s)")
	fmt.Println("[-] Waiting For", checkURL)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(checkURL)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("[!] %s List Download Fail!\n", target)
		fmt.Println("Exit AG-Installer 10s...!")
		time.Sleep(10 * time.Second)
		os.Exit(3)
	} else {
		fmt.Printf("[O] %s List Get Success!\n", target)
	}

	data, _ := io.ReadAll(resp.Body)
	agentInfos := strings.Split(string(data), "\r\n")
	agentInfo := strings.Split(agentInfos[0], " ")
	aName := agentInfo[0]
	aVer := agentInfo[1]

	return aName, aVer
}
