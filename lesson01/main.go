package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func task1() {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	log.Println("task1 done")
}

func task2() {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Println("task2 done")
}

func main() {
	wg.Add(2)

	go task1()
	go task2()

	wg.Wait()
}
