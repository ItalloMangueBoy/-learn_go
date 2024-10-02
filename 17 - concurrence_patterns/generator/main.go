package main

import (
	"fmt"
)

func write(txt string) <-chan string {
	txtCh := make(chan string)

	go func() {
		for {
			txtCh <- fmt.Sprintln(txt)
		}
	}()

	return txtCh
}

func main() {
	txtCh := write("Hello")

	for i := 0; i < 10; i++ {
		 fmt.Printf(<-txtCh)
	}
}
