package main

import (
	"fmt"
	"time"
)

func main() {
	stringChan := make(chan string, 1)
	intChan := make(chan int, 1)

	go func(){
		c := 0
		for{
			stringChan <- fmt.Sprintf("Hello %d", c)
			c++
			time.Sleep(2*time.Second)
		}
	}()

	go func(){
		c := 0
		for{
			intChan <- c
			c++
			time.Sleep(1*time.Second)
		}
	}()

	go func(){
		for{
			select{
			case v := <-stringChan:
				fmt.Println(v)
			case v := <-intChan:
				fmt.Println(v)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for{}
}
