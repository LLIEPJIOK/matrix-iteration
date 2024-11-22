package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/matrix-iteration/internal/application/iter"
)

func main() {
	if err := iter.Start(); err != nil {
		slog.Error(fmt.Sprintf("iter.Start(): %s", err))
		os.Exit(1)
	}
}
