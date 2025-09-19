package http

import (
	"fmt"
	"os"
)

var (
	DefaultWriter = os.Stdout
)

func FPrintln(format string, a ...any) {
	_, _ = fmt.Fprintf(DefaultWriter, format+"\n", a...)
}
