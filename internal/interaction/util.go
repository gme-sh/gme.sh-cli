package interaction

import "strings"

func orElse(c bool, a, b string) string {
	if c {
		return a
	} else {
		return b
	}
}

func repeat(str string, num int) string {
	var builder strings.Builder
	for i := 0; i < num; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
