package main

func findPrimeNumber(from uint32, to uint32, channel chan uint32) {

}

func main() {
	channel := make(chan uint32)

	go findPrimeNumber(1, 100, channel)
	go findPrimeNumber(100, 10000, channel)
	go findPrimeNumber(10000, 100000, channel)

	go func() {
		// Print every channel output in this goroutine
	}()
}
