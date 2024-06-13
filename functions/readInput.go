package functions

import "strings"

func ReadInput(input string, mp map[rune][]string) string {
	str := ""
	// clean := CheckInput(input)

	// fmt.Println(input)
	// words := strings.Split(input, " ")
	slInput := strings.Split(input, "\r\n")
	for _, word := range slInput {
		// i := 0
		for i := 0; i < 8; i++ {
			for _, char := range word {
				str += mp[char][i]
				// fmt.Print(line)
			}
			str += "\n"
			// i++
		}
	}

	return str
}
