// source/npm.go
package npm

import (
	"fmt"
	"os/exec"
	"registryhub/source"
	"registryhub/source/structs"
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

func (n NpmRegistryManager) SetRegistry(region structs.Region, sources *structs.RegistrySources) (string, error) {
	regionSources, ok := (*sources)[region]
	if !ok {
		return "", fmt.Errorf("unsupported region: %s", region)
	}

	npmSources, ok := regionSources["npm"]
	if !ok || len(npmSources) == 0 {
		return "", fmt.Errorf("npm sources not found for region: %s", region)
	}

	res := npmSources[0]

	c := exec.Command("npm", "config", "set", "registry", res)
	_, err := c.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return res, nil
}

func init() {
	source.RegisterManager([]string{"npm"}, NpmRegistryManager{})
}
