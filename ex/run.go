package ex

import (
	"os"
)

type Ex struct {
	Options
	ProgName string
}

type Options struct {
	IsVisual bool
	ReadOnly bool
	ShowMode bool
	Report   bool
	Magic    bool
}

func Run() error {
	var ex = Ex{}
	return ex.Main(os.Args)
}
