package main

import (
	"fmt"
)

func main(){
	channel := make(chan string, 2)

	channel <- "bob say"
	channel <- "alice say"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
