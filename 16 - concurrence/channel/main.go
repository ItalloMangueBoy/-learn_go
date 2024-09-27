package main

import (
	"fmt"
)

func write(txt string, txtChan chan string) {
	for count := 1; count <= 5; count++ {
		txtChan <- txt
	}

	close(txtChan)
}

func main() {
	var txtChan = make(chan string)

	go write("hello", txtChan)

	for msg := range txtChan {
		fmt.Println(msg)
	}
}
