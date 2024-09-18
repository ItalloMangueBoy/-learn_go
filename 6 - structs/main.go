package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	itallo := User{
		age:  20,
		name: "Itallo",
	}

	fmt.Println(itallo)
}
