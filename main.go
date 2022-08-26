package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}

// define a few variables.
var hunger = 3 // how many times a philosopher eats
// var eat = 1 * time.Second       // how long it takes to eat
// var think = 3 * time.Second     // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out
var orderFinished []string      // the order in which philosophers finish dining and leave
var orderMutex sync.Mutex       // a mutex for the slice orderFinished

// define a wait group
var wg sync.WaitGroup

func diningProblem(philosopher string, dominantHand, otherHand *sync.Mutex) {
	defer wg.Done()
	for i := hunger; i > 0; i-- {
		dominantHand.Lock()
		otherHand.Lock()

		dominantHand.Unlock()
		otherHand.Unlock()
	}

	// update the list of finished eaters
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}

func main() {
	fmt.Println("Dining Philosophers Problem")

	// add 5 (the number of philosophers) to the wait group
	wg.Add(len(philosophers))

	// we need to create a mutex for the very first fork (the one to
	// the left of the first philosopher). We create it as a pointer,
	// since a sync.Mutex must not be copied after its initial use.
	forkLeft := &sync.Mutex{}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the current philosopher's right fork, as a pointer
		forkRight := &sync.Mutex{}
		// fire off this philosopher's goroutine
		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}

	// wait for the philosophers to finish
	// this blocks until the wait group is 0
	wg.Wait()
	fmt.Printf("Order finished: %s\n", strings.Join(orderFinished, ", "))
	fmt.Println("The table is empty.")
}
