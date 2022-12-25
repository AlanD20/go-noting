package logger

import (
	"fmt"
	"strings"

	"github.com/labstack/gommon/color"
)

func Info(text ...string) {
	cyan := color.New().Cyan(strings.Join(text, ""))
	fmt.Println(cyan)
}

func Success(text ...string) {
	green := color.New().Green(strings.Join(text, ""))
	fmt.Println(green)
}

func Warn(text ...string) {
	yellow := color.New().Yellow(strings.Join(text, ""))
	fmt.Println(yellow)
}

func Error(text ...string) {
	red := color.New().Red(strings.Join(text, ""))
	fmt.Println(red)
}

func Panic(text ...string) {
	red := color.New().Red(strings.Join(text, ""))
	panic(red)
}
