package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Result struct {
	Result string
}

func worker(id int, jobChan <-chan string, resultChan chan<- Result, quitChan chan<- bool) {
	for job := range jobChan {
		log.Printf("worker %d started job %s", id, job)
		delay := rand.Intn(5) + 1
		time.Sleep(time.Duration(delay) * time.Second)

		resultChan <- Result{Result: fmt.Sprintf("worker %d finished job %s in %d seconds", id, job, delay)}
		log.Printf("worker %d finished job %s in %d seconds", id, job, delay)
	}

	quitChan <- true
}

func main() {
	jobChan := make(chan string, 20)
	resultChan := make(chan Result, 20)

	quitChan := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		go worker(i, jobChan, resultChan, quitChan)
	}

	for i := 0; i < 20; i++ {
		jobChan <- fmt.Sprintf("task-%d", i)
	}

	close(jobChan)

	closedWorkers := 0
	for {
		select {
		case result := <-resultChan:
			log.Println(result.Result)
		case <-quitChan:
			closedWorkers++
			if closedWorkers == 5 {
				log.Println("all workers closed")
				return
			}
		}
	}
}
