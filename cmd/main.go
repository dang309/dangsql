package main

import (
	"github.com/dang309/dangsql/internal/backend"
)

func main() {
	mb := backend.NewMemoryBackend()

	backend.RunRepl(mb)
}
