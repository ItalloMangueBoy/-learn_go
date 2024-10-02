package main

import "fmt"

func fibonacci(posição int) int {
	if posição <= 1 {
		return posição
	}

	return fibonacci(posição-2) + fibonacci(posição-1)

}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		fmt.Println(<-results)
	}
}
