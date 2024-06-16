// source/npm.go
package source

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

func (n NpmRegistryManager) SetRegistry(region string) (string, error) {
	rs, err := GetRemoteRegistrySources()
	if err != nil {
		return "", err
	}
	registry := (*rs)[StringToRegion(region)]["npm"][0]
	//TODO real set up registry
	err = exec.Command("npm", "config", "set", "registry", registry).Run()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return registry, nil
}
