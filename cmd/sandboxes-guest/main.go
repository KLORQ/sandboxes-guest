package main

import (
	"fmt"
	"os"

	"github.com/KLORQ/sandboxes-guest/internal/version"
)

func main() {
	_, _ = fmt.Fprintf(os.Stdout, "sandboxes-guest %s\n", version.Version)
}
