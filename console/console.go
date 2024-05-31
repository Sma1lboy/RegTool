package console

import (
	"fmt"
	"strings"
)

var Color = struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	White  string
}{
	Reset:  "\033[0m",
	Red:    "\033[31m",
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Blue:   "\033[34m",
	Purple: "\033[35m",
	Cyan:   "\033[36m",
	White:  "\033[37m",
}

func Println(color string, messages ...string) {
	message := strings.Join(messages, " ")
	fmt.Println(color + message + Color.Reset)
}
func Print(color string, messages ...string) {
	message := strings.Join(messages, " ")
	fmt.Print(color + message + Color.Reset)
}

func Printf(color string, format string, a ...interface{}) {
	fmt.Printf(color+format+Color.Reset, a...)
}

func Success(messages ...string) {
	Println(Color.Green, messages...)
}

func Error(messages ...string) {
	Println(Color.Red, messages...)
}

func Warning(messages ...string) {
	Println(Color.Yellow, messages...)
}

func Info(messages ...string) {
	Println(Color.Blue, messages...)
}
