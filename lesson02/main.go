package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func task1(data chan string) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	data <- "task1 done"
}

func task2(data chan string) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	data <- "task2 done"
}

func main() {
	data := make(chan string, 2)

	wg.Add(2)

	go task1(data)
	go task2(data)

	wg.Wait()

	close(data)

	for msg := range data {
		log.Println(msg)
	}
}
