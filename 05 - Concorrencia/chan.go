package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		time.Sleep(time.Second * 1)
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		time.Sleep(time.Millisecond * 400)
		c <- "pong"
	}
}



func main1() {

	c := make(chan string)

	go ping(c)
	go pong(c)

	go func () {
		for{
			msg := <- c
			fmt.Println(msg)
		}
	}()


	
	var teste string
	fmt.Scanf("%s" , &teste)

}
