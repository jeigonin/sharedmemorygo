package main

import "fmt"

func pong(pings <-chan int, pongs chan<- int, pings2 <-chan int, pongs2 chan<- int) {
	msg := <-pings
	msg++
	pongs <- msg

	msg2 := <-pings2
	msg2++
    pongs2 <- msg2
}


func main() {
    free_bed_channel := make(chan int, 1)
    used_bed_channel := make(chan int, 1)
	free_bed_channel <- 5 
	used_bed_channel <- 5 
    pong(free_bed_channel, free_bed_channel, used_bed_channel, used_bed_channel)
    fmt.Println(<-free_bed_channel)
    fmt.Println(<-used_bed_channel)
}