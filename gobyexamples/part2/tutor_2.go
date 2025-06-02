package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func _goroutines() {
	f("direct")
	// go keyword starts a new goroutine
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func _channels() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

func _channelBuffering() {
	// channel with buffer of size 2 to prevent deadlock
	msg := make(chan string, 2)

	msg <- "buffered"
	msg <- "channel"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func _channelSynchronization() {
	done := make(chan bool, 1)

	go worker(done)

	// if <-done is not used, the program will not wait for the worker to finish
	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func _channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

func _select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}

}

func _timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

func _nonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received.")
	}

	msg := "hola"
	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("No activity")
	}
}

func _closingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, isOpen := <-jobs
			if isOpen {
				fmt.Println("Received job: ", i)
			} else {
				fmt.Println("Received all jobs.")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job: ", i)
	}
	// time.Sleep(time.Second * 1)
	close(jobs)
	fmt.Println("Sent all jobs, Job closed.")

	<-done

	_, isOpen := <-jobs
	fmt.Println("received more jobs: ", isOpen)
}

func _rangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)
	q := <-queue
	// queue <- "three"
	fmt.Println(q)

	for elem := range queue {
		fmt.Println("Queue: ", elem)
	}
}

func main() {
	// Goroutines
	_goroutines()
	fmt.Println("---------------")
	// Channels
	_channels()
	fmt.Println("---------------")
	// Channel Buffering
	_channelBuffering()
	fmt.Println("---------------")
	// Channel Synchronization
	_channelSynchronization()
	fmt.Println("---------------")
	// Channel Directions
	_channelDirections()
	fmt.Println("---------------")
	// Select
	_select()
	fmt.Println("---------------")
	// Timeouts
	_timeouts()
	fmt.Println("---------------")
	// Non-Blocking Channel Operations
	_nonBlockingChannelOperations()
	fmt.Println("---------------")
	// Closing Channels
	_closingChannels()
	fmt.Println("---------------")
	// Range over Channels
	_rangeOverChannels()
	fmt.Println("---------------")
}
