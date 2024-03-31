package channels

import "fmt"

func RangeOverChannels1() {
	queue := make(chan string, 3)
	queue <- "One"
	queue <- "Two"
	queue <- "Three"

	close(queue)

	for ch := range queue {
		fmt.Println(ch)
	}

}

func RangeOverChannels2() {
	queue := make(chan string)

	go func() {
		defer close(queue)
		queue <- "One"
		queue <- "Two"
		queue <- "Three"
	}()

	for {
		ch, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(ch)
	}
}
