package main

import (
	"time"
	"fmt"
)

func main(){
	channel := make(chan int, 1)
	go func(channel chan int){
		channel <- 1
	}(channel)

	for{
		time.Sleep(1*time.Second)
		v := <- channel
		fmt.Println(v)
	}
}
