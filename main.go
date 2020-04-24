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


	// go change(free_bed_channel, used_bed_channel, 6, 4)
	// fmt.Println("free bed start", <-free_bed_channel)
	// fmt.Println("used bed start", <-used_bed_channel)

	no := 3
	var wg sync.WaitGroup
	count := 5
	count2 := 5
    for i := 0; i < no; i++ {
		count--
		count2++
        wg.Add(1)
		go change(free_bed_channel, used_bed_channel, count, count2, &wg)
		fmt.Println("\n")
		fmt.Println("free bed", <-free_bed_channel)
		fmt.Println("used bed", <-used_bed_channel)
    }
    wg.Wait()
}

// func doc1(bed chan<- int, bed2 chan<- int, wg *sync.WaitGroup){
// 	count--
// 	count2++
// 	wg.Add(1)
// 	change(free_bed_channel, used_bed_channel, count, count2, &wg)
	
// }

// func doc2(bed chan<- int, bed2 chan<- int, wg *sync.WaitGroup){
// 	count--
// 	count2++
// 	wg.Add(1)
// 	change(free_bed_channel, used_bed_channel, count, count2, &wg)

// }

func change(bed chan<- int, bed2 chan<- int, value int, value2 int,  wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	bed <- value
	bed2 <- value2
	time.Sleep(1 * time.Second)
	wg.Done()
}


// func take_bed(){
// 	free_bed_channel <-  -1 + <-free_bed_channel
// 	used_bed_channel <-  1+ <-used_bed_channel
// 	display_free_and_used_bed()
// }

// func display_free_and_used_bed(){
// 	fmt.Println("free bed updated", <-free_bed_channel)
// 	fmt.Println("used bed updated", <-used_bed_channel)
// }

// func free_bed(){
// 	free_bed_channel <- free_bed_channel +1
// 	used_bed_channel <- used_bed_channel -1
// 	display_free_and_used_bed()
// }