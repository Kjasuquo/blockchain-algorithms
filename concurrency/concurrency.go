package main

import (
	"sync"
)

func findPrimeNumber(from uint32, to uint32, channel chan uint32, wg *sync.WaitGroup) {

}

func main() {
	channel := make(chan uint32)
	var wg sync.WaitGroup

	go findPrimeNumber(1, 100, channel, &wg)
	go findPrimeNumber(100, 10000, channel, &wg)
	go findPrimeNumber(10000, 100000, channel, &wg)

	go func() {
		// Print every channel output in this goroutine
	}()
}
