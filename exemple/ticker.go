package main

var global int = 0
var c = make(chan int, 1)

func thread1(){
    <-c // Grab the ticket
    global = 1
    c <- 1 // Give it back
}

func thread2(){
    <-c
    global = 2
    c <- 1
}

func main() {
   c <- 1 // Put the initial value into the channel
   go thread1()
   go thread2()
}