package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func printAnimals() {
	var wg sync.WaitGroup
	var dogCnt, catCnt, fishCnt uint64
	dogChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	wg.Add(3)
	dogChan <- struct{}{}
	go printDog(&wg, dogChan, catChan, dogCnt)
	go printCat(&wg, catChan, fishChan, catCnt)
	go printFish(&wg, fishChan, dogChan, fishCnt)
	wg.Wait()
}

func printDog(wg *sync.WaitGroup, dogChan, catChan chan struct{}, counter uint64) {
	for {
		if counter > 100 {
			wg.Done()
			return
		}
		<-dogChan
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		catChan <- struct{}{}
	}
}

func printCat(wg *sync.WaitGroup, catChan, fishChan chan struct{}, counter uint64) {
	for {
		if counter > 100 {
			wg.Done()
			return
		}
		<-catChan
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		fishChan <- struct{}{}
	}
}

func printFish(wg *sync.WaitGroup, fishChan, dogChan chan struct{}, counter uint64) {
	for {
		if counter > 100 {
			wg.Done()
			return
		}
		<-fishChan
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		dogChan <- struct{}{}
	}
}

func main(){
	printAnimals()
}
