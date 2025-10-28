package main

import (
	"fmt"
	"sync"
)

type Job struct {
	JobId   int
	JobName string
}

func main() {
	wg := &sync.WaitGroup{}
	jobs := make(chan Job, 10)
	result := make(chan string, 10)

	// creating 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go WorkerPool(i, jobs, result, wg)
	}

	// assigining jobs
	for i := 1; i <= 10; i++ {
		jobs <- Job{JobId: i, JobName: "learning GO!!"}
	}
	close(jobs)

	wg.Wait()
	close(result)
	// collecting all the completed jobs
	for job := range result {
		fmt.Println(job)
	}
}

func WorkerPool(workerId int, jobs chan Job, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// job := <-jobs
	for job := range jobs {
		result <- fmt.Sprintf("worker id:%d processed jobid:%d job name:%v", workerId, job.JobId, job.JobName)
	}
}
