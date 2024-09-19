package main

import "fmt"

type People struct {
	name   string
	age    int
	height float64
	gender string
}

type Student struct {
	People
	curso    string
	facudade string
	periodo  int
}

func main() {
	itallo := Student{
		People: People{
			name:   "Itallo",
			age:    20,
			height: 1.7,
			gender: "masc",
		},
		curso:    "Analise e Desenvolvimento de Sistemas",
		facudade: "UNA",
		periodo:  3,
	}

	fmt.Println(itallo)
}
