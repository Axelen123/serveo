package internal

import (
	"fmt"
	"os"
)

// Error handles errors
func Error(prefix string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", prefix, err.Error())
	os.Exit(1)
}
