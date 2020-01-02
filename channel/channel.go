package main

import (
	"fmt"
	"log"
	"time"
)

func worker(done chan string, wait int) {
	fmt.Print("working...")
	time.Sleep(time.Duration(wait) * time.Second)
	fmt.Println("done")

	done <- "t"
}
func main() {

	done := make(chan string, 2)
	start := time.Now()
	go worker(done, 1)
	go worker(done, 2)

	<-done
	<-done
	elapsed := time.Since(start)
	log.Printf("Channel took %s", elapsed)
}
