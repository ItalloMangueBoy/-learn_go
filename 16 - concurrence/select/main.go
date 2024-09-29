package main

import (
	"fmt"
	"time"
)

func msgRoutine(msg string, sleep int) chan string {
	msgCh := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(sleep))
			msgCh <- msg
		}
	}()

	return msgCh	
}

func main() {
	ch1 := msgRoutine("1", 1)
	ch2 := msgRoutine("2", 5)

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)

		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
