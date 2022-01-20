package main

import "fmt"

func main() {
	fmt.Println("Resultado multiplicación =", multiplicacion(8, 10))
	fmt.Println("Resultado division =", division(10, 20))
}

func multiplicacion(x, y int) (multi int) {
	multi = x * y
	return
}

func division(x, y float64) (div float64) {
	div = x / y
	return
}
