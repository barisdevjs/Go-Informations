package channels

import "fmt"

func NonBlocking2() {
	messages := make(chan string, 2)
	signals := make(chan bool, 1)

	msg1 := "hi 1"
	select {
	case messages <- msg1:
		fmt.Println("Sent message", msg1)
	default:
		fmt.Println("No messages received")
	}

	msg2 := "hi 2"
	select {
	case messages <- msg2:
		fmt.Println("sent message", msg2)
	default:
		fmt.Println("no message sent")
	}

	boolFlag := true
	select {
	case signals <- boolFlag:
		fmt.Println("sent bool", boolFlag)
	default:
		fmt.Println("no bool sent")
	}

	select {
	case msg3 := <-messages:
		fmt.Println("received message", msg3)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
