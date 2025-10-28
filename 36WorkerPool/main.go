package main

import (
	"fmt"
	"time"
)

func main() {

	tasks := make(chan int, 10)
	result := make(chan string, 10)

	// created 5 workers
	for i := 1; i <= 5; i++ {
		go Worker(i, tasks, result)
	}

	// send 10 tasks
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	time.Sleep(1 * time.Second)
	close(result) // close result

	// collecting result
	for res := range result {
		fmt.Println(res)
	}
}

func Worker(worker int, tasks chan int, result chan string) {

	for task := range tasks {
		result <- fmt.Sprintf("worker %d processed task %d", worker, task)
	}
}
