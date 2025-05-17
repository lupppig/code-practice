package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 1; i <= 100; i += 2 {
			<-odd
			fmt.Printf("odd ==> %d ", i)
			even <- 1
		}
	}()

	go func() {
		for i := 0; i <= 100; i += 2 {
			<-even
			fmt.Printf("even ==> %d ", i)
			if i == 100 {
				done <- true
			}
			odd <- 1
		}
	}()

	even <- 1

	<-done
}
