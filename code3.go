package main

import (
	"fmt"
	"time"
)


func main(){
	channel := make(chan int, 1)

	channel <- 1

	for{
		go func(){
			v := <-channel
			fmt.Println(v)
		}()
		time.Sleep(1 * time.Second)
	}
}
