package main

import (
	"math"
)

// Definir 3 structs:
//Circulo, cuadrado, triangulo.

//Crear campos y metodos para calcular su area.

type Circulo struct {
	radio float64
}

type Cuadrado struct {
	lado float64
}

type Triangulo struct {
	base, altura float64
}

func (c Circulo) area() float64 {
	return math.Pow(c.radio, 2) * math.Pi
}

func (c Cuadrado) area() float64 {
	return math.Pow(c.lado, 2)
}

func (t Triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

func main() {

}
