package gem

import (
	"fmt"
	"os/exec"
	"regtool/common/alias"
	"regtool/source"
	"regtool/source/structs"
	"strings"
)

type GemRegistryManager struct{}

func (g GemRegistryManager) GetCurrRegistry() (string, error) {
	cmd := exec.Command("gem", "sources", "-l")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "https://") {
			return strings.TrimSpace(line), nil
		}
	}

	return "", fmt.Errorf("no valid source found")
}

func (g GemRegistryManager) SetRegistry(region structs.Region, sources *structs.RegistrySources) (string, error) {
	if sources == nil {
		return "", fmt.Errorf("sources is nil")
	}
	regionSources, ok := (*sources)[region]
	if !ok {
		return "", fmt.Errorf("unsupported region: %s", region)
	}
	source, ok := regionSources["gem"]
	if !ok || len(source) == 0 {
		return "", fmt.Errorf("gem sources not found for region: %s", region)
	}
	newSource := source[0]

	// Get current sources
	currentSources, err := g.getCurrentSources()
	if err != nil {
		return "", fmt.Errorf("error getting current sources: %v", err)
	}

	// Remove all current sources
	for _, src := range currentSources {
		removeCmd := exec.Command("gem", "sources", "--remove", src)
		_, err := removeCmd.Output()
		if err != nil {
			fmt.Printf("Error removing source %s: %v\n", src, err)
			// Continue with other sources even if one fails
		}
	}

	// Add the new source
	addCmd := exec.Command("gem", "sources", "--add", newSource)
	_, err = addCmd.Output()
	if err != nil {
		fmt.Println("Error adding new source:", err)
		return "", err
	}

	return newSource, nil
}

func (g GemRegistryManager) IsExists() bool {
	_, err := exec.Command("gem", "sources", "-l").Output()
	return err == nil
}

func (g GemRegistryManager) getCurrentSources() ([]string, error) {
	cmd := exec.Command("gem", "sources", "-l")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing gem sources -l: %v", err)
	}

	var sources []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "https://") {
			sources = append(sources, strings.TrimSpace(line))
		}
	}

	return sources, nil
}

func init() {
	alias.RegisterAlias("gem", []string{"rubygems"})
	source.RegisterManager([]string{"gem", "rubygems"}, GemRegistryManager{})
}
