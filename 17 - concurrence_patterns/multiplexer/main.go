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

func multiplexer(ch1, ch2 <-chan string) <-chan string {
	mainCh := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch1:
				mainCh <- msg

			case msg := <-ch2:
				mainCh <- msg
			}
		}
	}()

	return mainCh
}

func main() {
	ch1 := write("Hello")
	ch2 := write("world")

	txtCh := multiplexer(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Printf(<-txtCh)
	}
}
