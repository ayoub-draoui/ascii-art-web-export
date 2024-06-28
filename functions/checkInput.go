package functions

func CheckInput(input string) bool {
	isValid := true
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			if i < len(input)-1 && input[i] == '\r' && input[i+1] == '\n' {
				i++
				continue
			}
			isValid = false
			break
		}
	}
	return isValid
}
