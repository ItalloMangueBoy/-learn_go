package main

import "fmt"

func main()  {
	var v1 int = 1

	v2 := 2

	var (
		v3 int = 3
		v4 int = 4
	)

	v5, v6 := 5, 6

	fmt.Println(v1, v2, v3, v4, v5, v6)
}