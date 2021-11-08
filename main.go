package main

import (
	"fmt"
	"time"
)

func worker(n int, jobs chan int, result chan int) {
	fmt.Println("here", n)
	for j := range jobs {
		fmt.Println("Worker ", n, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker ", n, "finished job ", j)
		result <- j * 2
	}
}

const (
	noOfWorkers = 3
	noOfJobs    = 5
)

func main() {
	jobs := make(chan int, noOfJobs)
	result := make(chan int, noOfJobs)
	for i := 0; i < noOfWorkers; i++ {
		go worker(i, jobs, result)
	}

	for i := 1; i <= noOfJobs; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < noOfJobs; i++ {
		fmt.Println("result", <-result)
	}

}
