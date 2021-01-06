package main

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	num int
}

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan result, 1)
	commuChan := make(chan int, 1)

	go func() {
		// var a [10]result
		var a []result
		// var i = 0
		// re := a[0:]
		// result := make([]result, 0)
		for {
			r, ok := <-resultChan
			if ok {
				// a[i] = r
				// i++
				a = append(a, r)
			} else {
				break
			}
		}
		fmt.Println("close", a)
		commuChan <- 0
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	<-commuChan
	// time.Sleep(10 * time.Second)
}

func worker(i int, c chan result, wg *sync.WaitGroup) {
	// fmt.Println("i", i)

	defer wg.Done()
	time.Sleep(5 * time.Second)
	// fmt.Println("i", i)
	r := result{
		num: i,
	}
	c <- r
}
