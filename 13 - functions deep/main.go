package main

import (
	"fmt"
)

func Sum(numbers ...float64) (total float64) {
	for _, number := range numbers {
		total += number
	}

	return
}

func Sub(numbers ...float64) (total float64) {
	total = numbers[0]

	for _, number := range numbers[1:] {
		total -= number
	}

	return
}

func Mean(numbers ...float64) (mean float64) {
	if len(numbers) == 0 {
		return
	}

	sum := Sum(numbers...)
	count := float64(len(numbers))

	mean = sum / count

	return sum / count
}

func Calc(numbers ...float64) (sum float64, sub float64) {
	sum = Sum(numbers...)
	sub = Sub(numbers...)

	return
}

func IsApproved(scores ...float64) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if mean := Mean(scores...); mean > 6 {
		return true
	} else if mean < 6 {
		return false
	}

	panic("A media é exatamente 6")
}

func invertSignal(n *float64) {
	*n = *n * -1.0 
}

func init()  {
	println("Iniciando...")
}

func main() {
	defer fmt.Println("Observe os resultados acima")

	fmt.Println(IsApproved(6))
}
