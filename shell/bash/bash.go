package bash

import "registryhub/shell"

type Bash struct{}

func (b Bash) SetEnv(key, value string) error {
	return shell.SetEnvVarToFile(".bashrc", key, value)
}

func (b Bash) GetEnv(key string) (string, error) {
	return shell.GetEnvVarFromFile(".bashrc", key)
}

// init registers the Bash shell manager.
func init() {
	shell.RegisterShell("bash", func() shell.ShellManager {
		return Bash{}
	})
}
