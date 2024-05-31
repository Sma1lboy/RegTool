// source/docker.go
package source

import (
	"fmt"
	"os/exec"
)

func getDockerRegistry() string {
	// Replace with the actual command to get Docker registry
	cmd := exec.Command("docker", "info", "--format", "{{.RegistryConfig.IndexConfigs}}")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(output)
}
