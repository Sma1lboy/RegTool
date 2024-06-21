package shellenv

import (
	"fmt"
	"os"
	"registryhub/shellenv/bash"
	"registryhub/shellenv/zsh"
)

// SetEnv sets the environment variable for all supported shells
func SetEnv(key, value string) error {
	// set the environment variable for the current process
	err := os.Setenv(key, value)
	if err != nil {
		return fmt.Errorf("error setting %s: %w", key, err)
	}

	// create instances of all supported shells
	shells := []Shell{
		bash.Bash{},
		zsh.Zsh{},
	}

	// iterate over each shell and set the environment variable
	for _, shell := range shells {
		err := shell.SetEnv(key, value)
		if err != nil {
			return fmt.Errorf("error setting env in shell: %w", err)
		}
	}

	return nil
}

// GetEnv retrieves the value of the environment variable
func GetEnv(key string) (string, error) {
	// get the value of the environment variable from the current process
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s not set", key)
	}

	return value, nil
}
