package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	// test for termbox.

	termbox.SetCell(0, 0, 'あ', termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
	_ = termbox.PollEvent()

	w, h := termbox.Size()
	buf := termbox.CellBuffer()
	c0 := buf[0]
	c1 := buf[1]
	c2 := buf[2]
	fmt.Printf("w=%d h=%d\n", w, h)
	fmt.Printf("len(buf)=%d\n", len(buf))
	fmt.Printf("c0=%v\n", c0)
	fmt.Printf("c1=%v\n", c1)
	fmt.Printf("c2=%v\n", c2)
}
