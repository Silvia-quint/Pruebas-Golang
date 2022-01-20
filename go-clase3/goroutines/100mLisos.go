package main

import (
	"fmt"
	"sync"
)

type runner struct {
	name  string
	speed float64
}

func (r runner) checkName(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("I'm %s\n", r.name)
}

func main() {
	var wg sync.WaitGroup
	runners := [8]runner{
		{"Usain Bolt", 10.43841336},
		{"Tyson Gay", 10.31991744},
		{"Asafa Powell", 10.28806584},
		{"Justin Gatlin", 10.26694045},
		{"Christian Coleman", 10.24590164},
		{"Nesta Carter", 10.22494888},
		{"Maurice Greene", 10.2145046},
		{"Yohan Blake", 10.31991744},
	}
	// OTRA FORMA DE DECLARAR CADA CORREDOR -->
	// runners[0] = runner{name: "Usain Bolt", speed: 10.44}

	for i := 0; i < len(runners); i++ {
		wg.Add(1)
		go runners[i].checkName(&wg)
	}
	wg.Wait()
}
