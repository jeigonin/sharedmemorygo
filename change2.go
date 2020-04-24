package main

import(
	"fmt"
	"sync"
    "time"
)

var count int = 5
var count2 int = 5

func main() {
	free_bed_channel := make(chan int, 2)
	used_bed_channel := make(chan int, 2)
	free_bed_channel <- count
	used_bed_channel <- count2
	fmt.Println("free bed start", <-free_bed_channel)
	fmt.Println("used bed start", <-used_bed_channel)


	var wg sync.WaitGroup
	
	wg.Add(2)
	go doc1(free_bed_channel, used_bed_channel, free_bed_channel, used_bed_channel, &wg)
	fmt.Println("\n")
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)

	go doc2(free_bed_channel, used_bed_channel, free_bed_channel, used_bed_channel, &wg)
	fmt.Println("\n")
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)
	wg.Wait()
	
}

func doc1(read_free_bed_channel <-chan int, read_used_bed_channel <-chan int, write_free_bed_channel chan<- int, write_used_bed_channel chan<- int,wg *sync.WaitGroup){
	val1 := <-read_free_bed_channel + 1
	val2 := <-read_used_bed_channel - 1
	
	write(write_free_bed_channel, write_used_bed_channel, val1, val2)
	wg.Done()
}

func doc2(read_free_bed_channel <-chan int, read_used_bed_channel <-chan int, write_free_bed_channel chan<- int, write_used_bed_channel chan<- int,	wg *sync.WaitGroup){

	val1 := <-read_free_bed_channel + 1
	val2 := <-read_used_bed_channel - 1
	write(write_free_bed_channel, write_used_bed_channel, val1, val2)
	wg.Done()
}

func write(write_channel chan<- int, write_channel2 chan<- int, value int, value2 int) {
	time.Sleep(1 * time.Second)
	write_channel <- value
	write_channel2 <- value2
}


// select {
// case res := <-free_bed_channel:
// 	fmt.Println("free bed change ", res)
// case res := <-used_bed_channel:
// 	fmt.Println("used bed change", res)
// }