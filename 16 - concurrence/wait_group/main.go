package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for count := 0; count <= 5; count++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("hello")
		
		waitGroup.Done()
	}()

	go func() {
		write("world!!!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
