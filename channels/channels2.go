package channels

import "fmt"

func Channels2() {
	messages := make(chan string, 2)

	messages <- "1"
	messages <- "2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
