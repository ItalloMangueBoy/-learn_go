package main

import "fmt"

func sum(n1, n2 int) int {
	return n1 + n2
}

func sub(n1, n2 int) int {
	return n1 - n2
}

func mult(n1, n2 int) int {
	return n1 * n2
}

func div(n1, n2 int) int {
	return n1 / n2
}

func calculator(n1, n2 int) (int, int, int, int) {
	sum := sum(n1, n2)
	sub := sub(n1, n2)
	div := div(n1, n2)
	mult := mult(n1, n2)

	return sum, sub, div, mult
}

func main() {
	fmt.Println(calculator(10, 5))
}
