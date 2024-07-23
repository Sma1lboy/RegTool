package yarn

import (
	"fmt"
	"os/exec"
	"regtool/common/alias"
	"regtool/source"
	"regtool/source/structs"
	"strings"
)

type YarnRegistryManager struct{}

func (n YarnRegistryManager) GetCurrRegistry() (string, error) {
	cmd := exec.Command("yarn", "config", "get", "registry")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (n YarnRegistryManager) SetRegistry(region structs.Region, sources *structs.RegistrySources) (string, error) {
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

	fmt.Println(res)
	c := exec.Command("yarn", "config", "set", "registry", res)
	_, err := c.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return res, nil
}
func (n YarnRegistryManager) IsExists() bool {

	_, err := exec.Command("yarn", "config", "get", "registry").Output()

	return err == nil
}

func init() {
	alias.RegisterAlias("yarn", []string{})
	source.RegisterManager([]string{"yarn"}, YarnRegistryManager{})
}
