// source/npm.go
package source

import (
	"fmt"
	"os/exec"
	"strings"
)

func getNpmRegistry() string {
	cmd := exec.Command("npm", "config", "get", "registry")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}
