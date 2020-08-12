package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("caller started...")
	callerChan := make(chan string)
	hosts := os.Args[1:]

	for _, host := range hosts {
		go measure(host, callerChan)
	}

	for range hosts {
		fmt.Println(<-callerChan)
	}
}

func measure(host string, callerChan chan string) {
	startTime := time.Now()
	_, err := call(host)

	if err != nil {
		callerChan <- "failed to call " + host
		return
	}

	duration := time.Since(startTime).Seconds()
	callerChan <- fmt.Sprintf("%.2fs - %s", duration, host)
}

func call(host string) ([]byte, error) {
	response, err := http.Get(host)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
