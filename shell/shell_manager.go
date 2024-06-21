package shell

import (
	"fmt"
	"os"
	"strings"
)

// ShellManager is an interface for setting and getting environment variables in shell configuration files.
type ShellManager interface {
	SetEnv(key, value string) error
	GetEnv(key string) (string, error)
}

// ShellFactory is a function type that returns a new ShellManager.
type ShellFactory func() ShellManager

// shellFactories holds a map of shell names to their corresponding factory functions.
var shellFactories = map[string]ShellFactory{}

// RegisterShell registers a new shell factory.
func RegisterShell(name string, factory ShellFactory) {
	shellFactories[name] = factory
}

// NewShellManager creates a new ShellManager based on the current shell.
func NewShellManager() (ShellManager, error) {
	shell := os.Getenv("SHELL")
	for name, factory := range shellFactories {
		if strings.Contains(shell, name) {
			return factory(), nil
		}
	}
	return nil, fmt.Errorf("unsupported shell: %s", shell)
}
