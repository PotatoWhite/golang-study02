package main

import (
	"log"
	"time"
)

func task1(data chan string) {
	time.Sleep(5 * time.Second)
	data <- "task1 done"
}

func task2(data chan string) {
	time.Sleep(2 * time.Second)
	data <- "task2 done"
}

func main() {
	data := make(chan string, 2)

	go task1(data)
	go task2(data)

	for i := 0; i < 2; i++ {
		msg := <-data
		log.Println(msg)
	}
}
