package ex

import (
	"fmt"
	"path/filepath"
)

func TailPath(s string) string {
	return filepath.Base(s)
}

func (ex *Ex) mainInit(av []string) error {
	ex.ErrorPath = EXSTRINGS
	// Defend against d's, v's, w's, and a's in directories of
	// path leading to our true name.
	av0 := TailPath(av[0])

	// Figure out how we were invoked: ex, edit, vi, view.
	ex.Options.IsVisual = Any('v', av0)
	ex.Options.ReadOnly = Any('w', av0)
	if Any('d', av0) {
		ex.Options.ShowMode = true
		ex.Options.Report = true
		ex.Options.Magic = false
	}

	ex.ProgName = av0

	// TODO:

	return nil
}

func (ex *Ex) Main(av []string) error {

	err := ex.mainInit(av)
	if err != nil {
		return err
	}

	// TODO:
	fmt.Printf("%#v\n", ex)

	return nil
}
