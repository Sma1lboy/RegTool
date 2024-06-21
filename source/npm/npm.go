// source/npm.go
package npm

import (
	"fmt"
	"os/exec"
	"strings"
)

type NpmRegistryManager struct{}

func (n NpmRegistryManager) GetCurrRegistry() (string, error) {
	cmd := exec.Command("npm", "config", "get", "registry")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (n NpmRegistryManager) SetRegistry(registry string) (string, error) {
	err := exec.Command("npm", "config", "set", "registry", registry).Run()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return registry, nil
}
