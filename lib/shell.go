package lib

import (
	"bytes"
	"os"
	"os/exec"
)

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = append(os.Environ(),
		"GOPATH=/go",
		"GOLANG_VERSION=1.16.1",
	)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
