package zsh

import "registryhub/shell"

// Zsh represents the zsh shell.
type Zsh struct{}

// SetEnv writes the environment variable to the zsh configuration file.
func (z Zsh) SetEnv(key, value string) error {
	return shell.SetEnvVarToFile(".zshrc", key, value)
}

// GetEnv reads the environment variable from the zsh configuration file.
func (z Zsh) GetEnv(key string) (string, error) {
	return shell.GetEnvVarFromFile(".zshrc", key)
}

// init registers the Zsh shell manager.
func init() {
	shell.RegisterShell("zsh", func() shell.ShellManager {
		return Zsh{}
	})
}
