package helper

import (
	"os/exec"
	"strings"
)

func GenerateUUID() string {
	uuid, err := exec.Command("uuidgen").Output()
	PanicIfError(err)
	return strings.ToLower(strings.TrimSuffix(string(uuid), "\n"))
}
