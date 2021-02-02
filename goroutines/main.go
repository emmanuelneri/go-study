package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("hello")
		return
	}()
	wg.Wait()

	fmt.Println("world")

	goroutineResult()
	goroutineWaitAll()
}

func goroutineResult() {
	done := make(chan string)
	fmt.Println("call goroutine")
	go doSomeThing("hello", done)
	fmt.Println(<-done)
}

func doSomeThing(text string, done chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println(text)
	done <- "done"
}

func goroutineWaitAll() {
	var wg sync.WaitGroup
	quantity := 20
	fmt.Println("starting batch: ", quantity)
	for i := 0; i < quantity; i++ {
		wg.Add(1)
		go do(&wg, i)
	}

	fmt.Println("Waiting for batchs")
	wg.Wait()
	fmt.Println("batch completed")
}

func do(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("do %v: Started\n", id)
	time.Sleep(time.Second)
}
