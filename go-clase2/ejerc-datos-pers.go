package main

import "fmt"

// Crear un programa en el que dado un número de días, nos diga cuantas horas, minutos y segundos son.

// Usar tipos de datos personalizados.

type Dia uint
type Hora uint
type Minuto uint
type Segundo uint

func (d Dia) Hora() Hora {
	return Hora(d * 24)
}

func (h Hora) Minuto() Minuto {
	return Minuto(h * 60)
}

func (m Minuto) Segundo() Segundo {
	return Segundo(m * 60)
}

func main() {
	var d Dia

	d = 5
	cH := d.Hora()
	cM := cH.Minuto()

	fmt.Printf("Tenemos %d dias\n", d)
	fmt.Printf("%d dias son %d horas\n", d, cH /* cH = d.Hora() */)
	fmt.Printf("%d dias son %d minutos\n", d, cH.Minuto())
	fmt.Printf("%d dias son %d segundos\n", d, cM.Segundo())
}
