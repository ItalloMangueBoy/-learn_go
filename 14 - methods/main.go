package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (user User) presents() {
	fmt.Printf("Hello my name is %s!\n", user.Name)
}

func main() {
	itallo := User{
		Name: "Itallo",
		Age:  20,
	}

	itallo.presents()
}
