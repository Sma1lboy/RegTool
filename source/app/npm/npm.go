// source/npm.go
package npm

import (
	"fmt"
	"os/exec"
	"regtool/common/alias"
	"regtool/source"
	"regtool/source/structs"
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
	if sources == nil {
		return "", fmt.Errorf("sources is nil")
	}
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
func (n NpmRegistryManager) IsExists() bool {

	_, err := exec.Command("npm", "config", "get", "registry").Output()

	return err == nil
}

func init() {
	alias.RegisterAlias("npm", []string{})
	source.RegisterManager([]string{"npm"}, NpmRegistryManager{})
}
