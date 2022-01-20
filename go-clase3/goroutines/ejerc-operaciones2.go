package main

import (
	"fmt"
	"sync"
)

func suma(a, b int, wg *sync.WaitGroup, resultS *int) {
	defer wg.Done()
	*resultS = a + b
}

func multiplicacion(x, y int, wg *sync.WaitGroup, resultM *int) {
	defer wg.Done()
	*resultM = x * y

}

func division(x, y float64, wg *sync.WaitGroup, resultD *float64) {
	defer wg.Done()
	*resultD = x / y
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // wg.Add(1)
	var resultS int
	go suma(4, 7, &wg, &resultS)

	// wg.Add(1)
	var resultM int
	go multiplicacion(8, 10, &wg, &resultM)

	// wg.Add(1)
	var resultD float64
	go division(10, 20, &wg, &resultD)

	wg.Wait()
	fmt.Println("Resultado suma =", resultS)

	// wg.Wait()
	fmt.Println("Resultado multiplicación =", resultM)

	// wg.Wait()
	fmt.Println("Resultado division =", resultD)
}
