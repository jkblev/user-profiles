package main

import "fmt"

// goroutine that creates a dedicated channel
// that it can listen to for a stop signal. If the
// stop channel receives a signal, the goroutine immediately
// returns without finishing its loop.
func fibonacci(numInts int) (chan int, chan bool) {
	stop := make(chan bool)
	channel := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < numInts; i++ {
			select {
			case channel <- x:
				x, y = y, x+y
			case <-stop:
				return
			}
		}
	}()
	return channel, stop
}

func main() {
	number, stop := fibonacci(10)

	for i := 0; i < 5; i++ {
		fmt.Println(<-number)
	}
	stop <- true // stops the goroutine
}
