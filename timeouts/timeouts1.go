package timeouts

import (
	"fmt"
	"time"
)

func Timeouts1() {
	c1 := make(chan string, 1)
	// goroutine sends a value "R1" to channel c1 after sleeping for 2 seconds.
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "R1"
	}()

	select {
	// If a value is received from c1 within 1 second, it prints the received value.
	case res := <-c1:
		fmt.Println(res)
	// If no value is received from c1 within 1 second, it prints "Timeout 1".
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}
