package bash

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

// Bash represents the bash shell
type Bash struct{}

// SetEnv writes the environment variable to the bash configuration file
func (b Bash) SetEnv(key, value string) error {
	// get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}
	homeDir := usr.HomeDir

	// bash configuration file path
	filePath := fmt.Sprintf("%s/.bashrc", homeDir)

	return writeEnvVarToFile(filePath, key, value)
}

// writeEnvVarToFile writes the environment variable to the specified file
func writeEnvVarToFile(filename, key, value string) error {
	// read file content
	input, err := os.ReadFile(filename)
	if err != nil {
		// if the file does not exist, create a new file
		if os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()
		} else {
			return err
		}
	}

	lines := strings.Split(string(input), "\n")
	var output []string
	var found bool

	// update existing environment variable
	for _, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("export %s=", key)) {
			output = append(output, fmt.Sprintf("export %s=\"%s\"", key, value))
			found = true
		} else {
			output = append(output, line)
		}
	}

	// if not found, append new environment variable
	if !found {
		output = append(output, fmt.Sprintf("export %s=\"%s\"", key, value))
	}

	return os.WriteFile(filename, []byte(strings.Join(output, "\n")), 0644)
}
