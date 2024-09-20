package main

import "fmt"

func ifItallo() {
	var msg string

	if name := "Itallo"; name != "Itallo" {
		msg = "Hello normal people"
	} else if name == "Jenny" {
		msg = "Hello irmanzinha"
	} else {
		msg = "Hello gostosão"
	}

	fmt.Println(msg)
}

func switchDay(day int) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta "
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return ""
	}
}
