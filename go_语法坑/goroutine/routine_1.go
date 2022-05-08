package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	testStrings := make([]string, 16)

	for i := 0; i < len(testStrings); i++{
		testStrings[i] = "test" + strconv.Itoa(i)
	}

	for _, t := range testStrings{
		go func(){
			wg.Add(1)
			fmt.Println(t)
			wg.Done()
		}()
	}

	wg.Wait()
}
