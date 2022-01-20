package main

import "fmt"

func main() {
	var nota int
	result := ""

	fmt.Scanf("%d", &nota)

	switch {
	case nota <= 4:
		result = "suspendido"
	case nota <= 10:
		result = "aprobado"
	}
	fmt.Println(result)
}
