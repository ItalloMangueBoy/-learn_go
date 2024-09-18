package main

import "fmt"

func aritimetics(x, y int) {
	sum := x + y
	sub := x - y
	mult := x * y
	div := x / y
	mod := x % y

	fmt.Println(sum, sub, mult, div, mod)
}

func assigns() {
	var global_var string = "O operador '=' é utilizado para definir variaveis de escopo global"
	local_var := "O operador ':=' é utilizado para definir variaveis de escopo local"

	fmt.Println(global_var, local_var)
}

func ralational(x, y int) {
	fmt.Println(
		x > y,
		x >= y,
		x < y,
		x <= y,
		x == y,
		x != y,
	)
}

func logical(x, y bool) {
	fmt.Println(
		x && y,
		x || y,
		!x,
	)
}

func unary(x int) {
	x++
	x--
	
	x += 1
	x -= 1
	x *= 1
	x /= 1
	x %= 1

	fmt.Println(x)
}

func main() {
	fmt.Println("hello")
}
