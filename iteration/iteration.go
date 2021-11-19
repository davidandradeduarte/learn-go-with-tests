package iteration

import "strings"

// Repeat returns character repeated N times.
func Repeat(char rune, times int) string {
	repeated := strings.Builder{}
	for i := 0; i < times; i++ {
		repeated.WriteRune(char)
	}
	return repeated.String()
}
