package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	friends := [...]string{"Vlad", "Jenny", "GuiGui"}

	for key, friend := range friends {
		fmt.Println((key + 1), ":", friend)
	}

	for {
		fmt.Println("8")
	}
}
