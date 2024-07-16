package main

import (
	"regtool/cmd"
	"regtool/env"
	_ "regtool/source/initall"
)

func main() {
	env.Init()
	cmd.Run()
}
