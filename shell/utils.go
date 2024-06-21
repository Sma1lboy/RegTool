package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

// SetEnvVarToFile writes the environment variable to the specified shell configuration file.
func SetEnvVarToFile(filename, key, value string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}
	homeDir := usr.HomeDir
	filePath := fmt.Sprintf("%s/%s", homeDir, filename)

	input, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)
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

	for _, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("export %s=", key)) {
			output = append(output, fmt.Sprintf("export %s=\"%s\"", key, value))
			found = true
		} else {
			output = append(output, line)
		}
	}

	if !found {
		output = append(output, fmt.Sprintf("export %s=\"%s\"", key, value))
	}

	return os.WriteFile(filePath, []byte(strings.Join(output, "\n")), 0644)
}

// GetEnvVarFromFile reads the environment variable from the specified shell configuration file.
func GetEnvVarFromFile(filename, key string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting current user: %w", err)
	}
	homeDir := usr.HomeDir
	filePath := fmt.Sprintf("%s/%s", homeDir, filename)

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, fmt.Sprintf("export %s=", key)) {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.Trim(parts[1], "\""), nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("environment variable %s not found", key)
}

// GetEnv attempts to get the environment variable from the OS environment first,
// and if not found, it reads from the specified shell configuration file.
func GetEnv(key, filename string) (string, error) {
	// First, try to get the environment variable from the OS environment
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	// If not found in OS environment, proceed to read from the shell configuration file
	return GetEnvVarFromFile(filename, key)
}
