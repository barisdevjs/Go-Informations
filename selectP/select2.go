package selectP

import (
	"fmt"
	"time"
)

type Message struct {
	Msg      string
	Received time.Time
}

func sendMessageWithDelay(delay time.Duration, message string, c chan<- Message) {
	time.Sleep(delay)
	c <- Message{message, time.Now()}
}

func Select2() {
	// Create channels and messages slice
	channels := []chan Message{
		make(chan Message),
		make(chan Message),
	}
	messages := []string{"one", "two"}

	// Start goroutines
	for i := range channels {
		go sendMessageWithDelay(time.Duration(i+1)*time.Second, messages[i], channels[i])
	}

	// Receive messages
	for range channels {
		select {
		case msg := <-channels[0]:
			fmt.Printf("Received from c1: %s at %s\n", msg.Msg, msg.Received.Format("15:04:05"))
			fmt.Printf("Time elapsed: %v\n", time.Since(msg.Received))
		case msg := <-channels[1]:
			fmt.Printf("Received from c2: %s at %s\n", msg.Msg, msg.Received.Format("15:04:05"))
			fmt.Printf("Time elapsed: %v\n", time.Since(msg.Received))
		}
	}
}
