package main

import "fmt"

func main() {
	slicePar := []int{}
	sliceImpar := []int{}
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			slicePar = append(slicePar, i)
		} else {
			sliceImpar = append(sliceImpar, i)
		}
	}

	fmt.Println("Numeros Pares: ", slicePar)
	fmt.Println("Numeros impares: ", sliceImpar)
}
