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
	fmt.Println("free bed start", <-free_bed_channel)
	fmt.Println("used bed start", <-used_bed_channel)

	var wg sync.WaitGroup
	
	wg.Add(2)
	go doc1(free_bed_channel, used_bed_channel, &wg)
	go doc2(free_bed_channel, used_bed_channel, &wg)
	
	wg.Wait()
    
}

func doc1(free_bed_channel <-chan int, used_bed_channel <-chan int, wg *sync.WaitGroup){
	
	change(free_bed_channel, used_bed_channel)
	wg.Done()
	
	fmt.Println("\n")
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)
}

func doc2(free_bed_channel <-chan int, used_bed_channel <-chan int, wg *sync.WaitGroup){
	change(free_bed_channel, used_bed_channel)
	wg.Done()
	
	fmt.Println("\n")
	fmt.Println("free bed", <-free_bed_channel)
	fmt.Println("used bed", <-used_bed_channel)
}

func change(free_bed_channel chan<- int, used_bed_channel chan<- int, value int, value2 int){

}

func use_bed(free_bed_channel <-chan int, used_bed_channel <-chan int) {
	time.Sleep(1 * time.Second)
	free_bed_channel <- ( -1 + <- free_bed_channel )
	used_bed_channel <- ( 1 + <- free_bed_channel )
}

func free_bed(free_bed_channel <-chan int, used_bed_channel <-chan int) {
	
	time.Sleep(1 * time.Second)
	free_bed_channel <- ( 1 + <- free_bed_channel )
	used_bed_channel <- ( -1 + <- free_bed_channel )
}
