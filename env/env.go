package env

import "os"

const (
	DEBUG = "DEBUG"
	PROD  = "PROD"
	GO_MODE = "GO_MODE"
)

func Init() {
	os.Setenv(GO_MODE, DEBUG)
}

