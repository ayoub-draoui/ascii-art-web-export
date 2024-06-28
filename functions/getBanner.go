package functions

import (
	"bufio"
	"fmt"
	"os"
)

func GetBanner(banner string) map[rune][]string {
	mp_banner := make(map[rune][]string)
	file, err := os.Open("./sources/" + banner + ".txt")
	if err != nil {
		fmt.Println("there is a problem opening the file")
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 31
	line := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			i++
			continue
		}
		mp_banner[rune(i)] = append(mp_banner[rune(i)], scanner.Text())
		line++
	}

	return mp_banner
}
