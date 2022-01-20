package main

type X struct {
	Titulo string
	Id
	Data []Data
}

type Subsubsubid struct {
	Value int `json:"value"`
}

type Subsubid struct {
	Subsubsubid `json:"subsubsubid"`
}

type Subid struct {
	Subsubid `json:"subsubid"`
}

type Id struct {
	Subid `json:"subid"`
}
type Product struct {
}
type Data struct {
	Product
	Activo bool
	Price  float64
}

func main() {

}
