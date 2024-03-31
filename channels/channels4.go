package channels

import "fmt"

func ping(pings chan<- string, msg string) { // sender
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // receive + send
	msg := <-pings
	pongs <- msg
}

func Channels4() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Pass to sender")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
