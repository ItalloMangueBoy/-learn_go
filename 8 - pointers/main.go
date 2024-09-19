package main

import "fmt"

func main() {
	num := 100
	numPointer := &num
	numOriginal := num

	num *= 2

	fmt.Println(num, numPointer, *numPointer, numOriginal)
}
