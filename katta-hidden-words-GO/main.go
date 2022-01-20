package main

import (
	"fmt"
	"strconv"
	"strings"
)

type code struct {
	letter string
}

var result = []string{}

func main() {
	var value int
	fmt.Println("Introduce un valor entre 100 y 999999:")
	fmt.Scanf("%d", &value)

	codes := []string{"o", "b", "l", "i", "e", "t", "a", "d", "n", "m"}
	//codes := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "ñ", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if value > 99 && value <= 999999 {
		codeNumber := strconv.Itoa(value)
		arrayNumber := strings.Split(codeNumber, "")

		for i := 0; i < len(arrayNumber); i++ {
			letterbyNumber := arrayNumber[i]
			number, _ := strconv.Atoi(letterbyNumber)
			x := codes[number]
			result = append(result, fmt.Sprintf("%s", x))
		}
		fmt.Println(strings.Join(result, ""))
	}
}
