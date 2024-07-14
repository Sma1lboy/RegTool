package main

import (
	"registryhub/cmd"
	"registryhub/env"
	_ "registryhub/source"
)

func main() {
	env.Init()
	cmd.Run()
}
