package di

import (
	"fmt"
	"io"
)

// Greet func
func Greet(w io.Writer, str string) {
	fmt.Fprintf(w, "Hello, %s!", str)
}
