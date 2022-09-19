package main

import (
	"AG-Installer/utils"
	"fmt"
	"os"
	"time"
)

var inis utils.Inis

func main() {
	fmt.Println("==========================")
	fmt.Println("Anti-Gravity Installer 1.0")
	fmt.Println("==========================")

	taskName := "AntiGravityUpdateService"
	taskName2 := "AntiGravityAgentService"

	// OS 종류 확인
	osv := utils.CheckOS()

	// ini 파일 로드
	utils.CheckIni("setting.ini", &inis)

	// 최신 Agent&Updater 리스트 획득
	aName, aVer := utils.CheckList("Agent", utils.GetAgentListUrl(&inis, osv, "agent"))
	aName2, aVer2 := utils.CheckList("Updater", utils.GetAgentListUrl(&inis, osv, "updater"))

	// 최신 Agent&updater 다운로드 주소 획득
	downURL := inis.ServerProtocol + "://" + inis.ServerIp + ":" + inis.ServerPort + "/" + osv + "/AG-Agent/" + aVer + "/" + aName
	downURL2 := inis.ServerProtocol + "://" + inis.ServerIp + ":" + inis.ServerPort + "/" + osv + "/AG-Updater/" + aVer2 + "/" + aName2

	// 운영체제에 맞게 다운로드한 파일에 저장
	switch osv {
	case "windows":
		// Windows의 기본 폴더 생성
		value := os.Getenv("ProgramFiles")

		err := os.Mkdir(value+"\\Anti-Gravity", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		// Agent&Updater 다운로드
		path, _ := utils.Downloads(downURL, value+"\\Anti-Gravity\\"+aName)
		path2, _ := utils.Downloads(downURL2, value+"\\Anti-Gravity\\"+aName)

		// Agent Service 등록
		_ = utils.AddService(taskName, path)
		_ = utils.AddService(taskName2, path2)

		// Updater Service  실행
		_ = utils.RunService(taskName2)
		_ = utils.RunService(taskName)

		fmt.Println("============설치 완료=============")
		time.Sleep(10 * time.Second)

	case "linux":
		osv = "linux"
	}

	// TODO
}
