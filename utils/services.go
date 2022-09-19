package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}
func printError(err error) {
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func RunService(taskName string) error {
	prog := "sc.exe"
	arguments := []string{}
	arguments = append(arguments, "start")
	arguments = append(arguments, taskName)

	time.Sleep(5 * time.Second)
	cmd2 := exec.Command(prog, arguments...)
	printCommand(cmd2)
	outupt, err := cmd2.CombinedOutput()
	printError(err)
	printOutput(outupt)
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
	printCommand(cmd)
	outupt, err := cmd.CombinedOutput()
	printError(err)
	printOutput(outupt)
	return err
}
