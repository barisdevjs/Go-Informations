package channels

import "fmt"

func CloseChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all Jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job", j)
	}

	close(jobs)
	fmt.Println("Sent ALL Jobs")
	<-done

	_, ok := <-jobs
	fmt.Println("Received more jobs", ok)
}
