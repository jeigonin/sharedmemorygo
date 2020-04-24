package main

import "fmt"
import "sync"

func doc1(free_bed_channel chan int, wg *sync.WaitGroup) {

	msg:= read(free_bed_channel)
	fmt.Println(<-free_bed_channel)
	msg++
	free_bed_channel <- msg

	wg.Done()
}

var wg sync.WaitGroup

func main() {
    free_bed_channel := make(chan int, 1)
	free_bed_channel <- 5 

    fmt.Println(<-free_bed_channel)
	wg.Add(1)
    go doc1(free_bed_channel, &wg)
	fmt.Println(<-free_bed_channel)
	
	wg.Wait()
}

func write(ch chan int, val int){
	ch <- val
}

func read(ch chan int){
	return <-ch
}