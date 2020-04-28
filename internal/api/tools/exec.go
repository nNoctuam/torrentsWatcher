package tools

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func ExecShell(command string, arguments ...string) string {
	output, err := TryExecShell(command, arguments...)
	if err != nil {
		fmt.Printf("Failed to execute [%s %s]\n", command, strings.Join(arguments, " "))
		fmt.Printf(output)
		panic(err)
	}
	return output
}

func TryExecShell(command string, arguments ...string) (string, error) {
	segments := strings.Fields(command)
	segments = append(segments, arguments...)
	//fmt.Printf("executing [%s %s]\n", segments[0], strings.Join(segments[1:], " "))
	cmd := exec.Command(segments[0], segments[1:]...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		return stderr.String(), err
	}
	return string(output), nil
}
