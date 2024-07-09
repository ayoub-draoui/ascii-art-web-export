package functions

func CheckInput(input string) string {
	// isValid := true
	var str string
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			if i < len(input)-1 && input[i] == '\r' || input[i] == '\n' {
				str += string(input[i])
			}
			// isValid = false
			// break
		} else {
			str += string(input[i])
		}
	}

	return str
}
