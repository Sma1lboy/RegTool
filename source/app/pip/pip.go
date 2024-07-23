package pip

import (
	"fmt"
	"os/exec"
	"regtool/common/alias"
	"regtool/source"
	"regtool/source/structs"
	"strings"
)

type PipRegistryManager struct{}

func (n PipRegistryManager) GetCurrRegistry() (string, error) {
	cmd := exec.Command("pip", "config", "get", "global.index-url")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (n PipRegistryManager) SetRegistry(region structs.Region, sources *structs.RegistrySources) (string, error) {
	if sources == nil {
		return "", fmt.Errorf("sources is nil")
	}
	regionSources, ok := (*sources)[region]
	if !ok {
		return "", fmt.Errorf("unsupported region: %s", region)
	}

	source, ok := regionSources["npm"]
	if !ok || len(source) == 0 {
		return "", fmt.Errorf("npm sources not found for region: %s", region)
	}

	res := source[0]

	fmt.Println(res)
	c := exec.Command("pip", "config", "set", "global.index-url", res)
	_, err := c.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return res, nil
}
func (n PipRegistryManager) IsExists() bool {

	_, err := exec.Command("pip", "config", "get", "global.index-url").Output()

	return err == nil
}

func init() {
	alias.RegisterAlias("pip", []string{})
	source.RegisterManager([]string{"pip"}, PipRegistryManager{})
}
