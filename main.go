package main

import (
	"math/rand"
	"sync"
)

const numberOfGoroutines = 100_000

// A global slice that is incorrectly shared among multiple goroutines
var globalSlice []string

func main() {
	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	for tries := 0; tries < numberOfGoroutines; tries++ {
		go func() {
			defer wg.Done()

			// Launching a new goroutine to modify the global slice
			// and sending the result back via a channel
			sliceChan := make(chan []string)
			go alterGlobalSlice(sliceChan)
			// Seems like the GC is not keeping track of `theSlice` on this line
			theSlice := <-sliceChan
			// If you clone the slice, the error won't happen
			// If you unblock the follow 2 lines, make sure you use theSlice2 in the next for loop
			//var theSlice2 []string
			//copy(theSlice2, theSlice)

			// Attempting to access elements by iterating through the slice
			someMap := make(map[string]string)
			for _, theSliceElement := range theSlice {
				someMap[theSliceElement] = theSliceElement
			}
		}()
	}

	wg.Wait()
}

func alterGlobalSlice(sliceChan chan []string) {
	globalSlice = make([]string, rand.Intn(15)+1)

	sliceChan <- globalSlice
}
