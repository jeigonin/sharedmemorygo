package main

import(
	"fmt"
	"sync"
	"time"
)


func main() {
	free_bed_channel := make(chan int, 2)
	used_bed_channel := make(chan int, 2)
	free_bed_channel <- 5
	used_bed_channel <- 5
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)


	var wg sync.WaitGroup
	
	wg.Add(2)
	
	go doc1(free_bed_channel, free_bed_channel, used_bed_channel, used_bed_channel, &wg)
	fmt.Println("\n")
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)

	go doc2(free_bed_channel, free_bed_channel, used_bed_channel, used_bed_channel, &wg)
	fmt.Println("\n")
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)
	wg.Wait()
	
}

func doc1(read_free_bed_channel <-chan int, write_free_bed_channel chan<- int, read_used_bed_channel <-chan int, write_used_bed_channel chan<- int, wg *sync.WaitGroup) {
	
	fmt.Println("doc1")
	msg := <-read_free_bed_channel
	msg++

	msg2 := <-read_used_bed_channel
	msg2--
	
	change(write_free_bed_channel, write_used_bed_channel, msg, msg2)

	wg.Done()
	
}

func doc2(read_free_bed_channel <-chan int, write_free_bed_channel chan<- int, read_used_bed_channel <-chan int, write_used_bed_channel chan<- int, wg *sync.WaitGroup) {
	fmt.Println("doc2")
	msg := <-read_free_bed_channel
	msg--
	
	msg2 := <-read_used_bed_channel
	msg2++
	
	change(write_free_bed_channel, write_used_bed_channel, msg, msg2)
	
	wg.Done()
}

func change(bed chan<- int, bed2 chan<- int, value int, value2 int) {
	fmt.Println("change")
    time.Sleep(1 * time.Second)
    bed <- value
    bed2 <- value2
}


// select {
// case res := <-free_bed_channel:
// 	fmt.Println("free bed change ", res)
// case res := <-used_bed_channel:
// 	fmt.Println("used bed change", res)
// }