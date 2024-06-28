package functions

import (
	"strings"
)

func ReadInput(input string, mp map[rune][]string) string {
	str := ""
	slInput := strings.Split(input, "\r\n")
	for _, word := range slInput {
		if word == "" {
			str += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					str += mp[char][i]
				}
				str += "\n"
			}
		}
	}

	return str
}
