package main

import (
	"registryhub/cmd"
	"registryhub/env"
	_ "registryhub/source/initall"
)

func main() {
	env.Init()
	cmd.Run()
}
