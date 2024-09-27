package main

import (
	"os"

	"github.com/kkentzo/sec2env/cmd"
)

func main() {
	root := cmd.New()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
