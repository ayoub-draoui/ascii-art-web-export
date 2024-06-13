package functions

import "fmt"

func CheckInput(input string) bool {
	isValid := true
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			fmt.Println("Please Enter Valid Input")
			isValid = false
			break
		}
	}
	return isValid
}
