package main

import (
	"fmt"

	"github.com/perdasilva/deconflictor/internal"
	"github.com/perdasilva/deconflictor/internal/samples"
	"github.com/perdasilva/deconflictor/internal/stylizer"
)

func main() {
	for _, sample := range samples.ErrorSamples {
		runner := internal.NewRunner(&stylizer.PrettyTree{}, sample)
		runner.Run()
		fmt.Printf("\n\n")
	}
}
