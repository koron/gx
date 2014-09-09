package ex

import (
	"strings"
)

func Any(c rune, s string) bool {
	return strings.IndexRune(s, c) >= 0
}
