package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}
func printError(err error, taskName string, path string) {
	if err != nil {
		if strings.Contains(err.Error(), "1073") {
			_ = DeleteService(taskName)
			_ = AddService(taskName, path)
			_, _ = fmt.Println("==> Service 갱신 등록", taskName)
		}
	}
}

func printOutput(outs []byte, taskName string) {
	if len(outs) > 0 {
		if strings.Contains(string(outs), "AntiGravityAgentService") || strings.Contains(string(outs), "AntiGravityUpdateService") {
			_, _ = fmt.Println("==> Service 시작", taskName)
		}
	}
}

func DeleteService(taskName string) error {
	prog := "sc.exe"
	arguments := []string{}
	arguments = append(arguments, "delete")
	arguments = append(arguments, taskName)

	cmd2 := exec.Command(prog, arguments...)

	outupt, err := cmd2.CombinedOutput()
	printError(err, taskName, "")
	printOutput(outupt, taskName)
	return err
}

func RunService(taskName string) error {
	prog := "sc.exe"
	arguments := []string{}
	arguments = append(arguments, "start")
	arguments = append(arguments, taskName)

	cmd2 := exec.Command(prog, arguments...)

	outupt, err := cmd2.CombinedOutput()
	printError(err, taskName, "")
	printOutput(outupt, taskName)
	return err
}

func AddService(taskName string, path string) error {

	prog := "sc.exe"
	arguments := []string{}
	arguments = append(arguments, "create")
	arguments = append(arguments, taskName)
	arguments = append(arguments, "binPath=")
	arguments = append(arguments, path)

	arguments = append(arguments, "start=")
	arguments = append(arguments, "auto")

	arguments = append(arguments, "DisplayName=")
	arguments = append(arguments, taskName)

	cmd := exec.Command(prog, arguments...)

	outupt, err := cmd.CombinedOutput()
	printError(err, taskName, path)
	printOutput(outupt, taskName)
	return err
}
