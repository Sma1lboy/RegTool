package main

import (
	"registryhub/cmd"
	"registryhub/env"
)

func main() {
	env.Init()
	cmd.Execute()
}
