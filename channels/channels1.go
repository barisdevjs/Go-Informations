package channels

import "fmt"

func Channel1() {
	messages := make(chan string)

	go func() { messages <- "Ping !" }()

	msg := <-messages
	fmt.Println(msg)
}
