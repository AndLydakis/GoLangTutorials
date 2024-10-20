package main

import (
	"fmt"
	"sync"
	"time"
)

func coroutines() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Async exec")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("sync exec")
}

func channels() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go func() {
		ch <- "the message"
	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

func select_stm() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "msg to channel 1"
	}()

	go func() {
		ch2 <- "msg to channel 2"
	}()

	time.Sleep(10 * time.Millisecond) // both are true, random select
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("none")
	}

}

func looping() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // so the loop below knows to end
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

}

func main() {
	// coroutines()
	// channels()
	// select_stm()
	looping()
}
