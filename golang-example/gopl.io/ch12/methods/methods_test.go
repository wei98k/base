package methods

import (
	"strings"
	"time"
)

func PrintDuration() {
	Print(time.Hour)
}

func ExamplePrintDuration() {
	PrintDuration()
}

func PrintReplacer() {
	Print(new(strings.Replacer))
}

func ExamplePrintReplacer() {
	PrintReplacer()
}
