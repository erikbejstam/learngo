package iteration

import "strings"

func Repeat(c string, n int) string {
	var repeated strings.Builder

	c = strings.Trim(c, " ?!.")
	c = strings.ToLower(c)

	for range n {
		repeated.WriteString(c)
	}
	return repeated.String()
}