package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetPath() (string, error) {
	// system username
	out, err := exec.Command("whoami").Output()
	if err != nil {
		return "", err
	}

	user := strings.TrimSpace(string(out))
	return fmt.Sprintf("/home/%s/.jenkins/", user), nil
}
