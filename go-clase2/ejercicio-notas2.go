package main

import "fmt"

func main() {
	var nota int
	result := ""

	fmt.Scanf("%d", &nota)

	switch {
	case nota <= 4:
		result = "suspenso"
	case nota <= 6:
		result = "aprobado"
	case nota <= 8:
		result = "notable"
	case nota <= 10:
		result = "sobresaliente"
	}
	fmt.Println(result)
}
