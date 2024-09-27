package main

import (
	"fmt"
	"time"
)

func write(text string) {
	for count := 0; count <= 5; count++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	go write("hello")
	write("world!!!")
}
