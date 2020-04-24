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
	go doc1(free_bed_channel, used_bed_channel, &wg)
    fmt.Println("\n")
    fmt.Println("free bed", <-free_bed_channel)
    fmt.Println("used bed", <-used_bed_channel)

    go doc2(free_bed_channel,free_bed_channel, used_bed_channel, &wg)
    fmt.Println("\n")
    fmt.Println("free bed", <-free_bed_channel)
    fmt.Println("used bed", <-used_bed_channel)
    wg.Wait()

}
func doc1(free_bed_channel chan<- int, used_bed_channel chan<- int, wg *sync.WaitGroup){

    count++
    count2--
    change(free_bed_channel, used_bed_channel, count, count2)
    wg.Done()
}

func doc2(free_bed_channel chan<- int, free_bed_channel2 <-chan int, used_bed_channel chan<- int, wg *sync.WaitGroup){
    count--
    count2++
    change(free_bed_channel, used_bed_channel, count, count2)
    wg.Done()
}

func change(bed chan<- int, bed2 chan<- int, value int, value2 int) {
	fmt.Println("change")

    time.Sleep(1 * time.Second)
    bed <- value
    bed2 <- value2
}