package main

import (
	"fmt"
	"log/slog"
	"os"

	"matrix-iter/internal/application/iter"
)

func main() {
	if err := iter.Start(); err != nil {
		slog.Error(fmt.Sprintf("iter.Start(): %s", err))
		os.Exit(1)
	}
}
