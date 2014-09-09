package main

import (
	"github.com/koron/gx/ex"
	"os"
)

func main() {
	err := ex.Run()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
