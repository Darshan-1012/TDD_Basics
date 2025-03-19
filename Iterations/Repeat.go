package iterations

import "strings"

const repeatednum = 5

func repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatednum; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
