package bash

import "registryhub/shell"

// Bash represents the bash shell.
type Bash struct{}

// SetEnv writes the environment variable to the bash configuration file.
func (b Bash) SetEnv(key, value string) error {
	return shell.SetEnvVarToFile(".bashrc", key, value)
}

// GetEnv reads the environment variable from the bash configuration file.
func (b Bash) GetEnv(key string) (string, error) {
	return shell.GetEnvVarFromFile(".bashrc", key)
}

// init registers the Bash shell manager.
func init() {
	shell.RegisterShell("bash", func() shell.ShellManager {
		return Bash{}
	})
}
