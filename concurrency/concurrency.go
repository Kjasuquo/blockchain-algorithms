package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup

func findPrimeNumber(from uint32, to uint32, channel chan uint32) {
	wg.Add(1)
	if to < 2 {
		return
	}
	for from <= to {

		isPrime := true
		if from == 1 {
			isPrime = false
		}
		sqRoot := math.Sqrt(float64(from))
		for i := 2; i <= int(sqRoot); i++ {
			if int(from)%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			channel <- from
		}
		from++

	}
	wg.Done()
	wg.Wait()

}

func main() {

	// Keeping track of time taken to execute the function
	startTime := time.Now()

	wg.Add(1)

	channel := make(chan uint32)

	go findPrimeNumber(1, 100, channel)
	go findPrimeNumber(100, 10000, channel)
	go findPrimeNumber(10000, 100000, channel)

	go func() {

		//Print every channel output in this goroutine

		for {

			val, ok := <-channel
			fmt.Println(val)
			if !ok {
				close(channel)
				break
			}

		}

	}()

	// Time duration
	TimeDuration := time.Since(startTime)
	fmt.Println("Time Duration", TimeDuration)

	wg.Done()
	wg.Wait()

}
