package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	} 

	fmt.Println("f1 execution finished")

	//3.
    // Before exiting, call wg.Done() in each goroutine
    // to indicate to the WaitGroup that the goroutine has finished executing.
    wg.Done()
    //or:
    // (*wg).Done()
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")

	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}

	fmt.Println("f2 execution finished")
}

func main() {
	fmt.Println("main execution started")
	// 1.
    // Create a new instance of sync.WaitGroup (we’ll call it symply wg)
    // This WaitGroup is used to wait for all the goroutines that have been launched to finish.
	var wg sync.WaitGroup
 	// 2.
    // Call wg.Add(n) method before attempting to
    // launch the go routine.
	wg.Add(1)
	// Launching a goroutine
	go f1(&wg)
  	// No. of running goroutines
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())
	// calling other functions:
	f2()
	// time.Sleep(time.Second * 2)
	// 4.
    // Finally, we call wg.Wait()to block the execution of main() until the goroutines
    // in the WaitGroup have successfully completed.
	wg.Wait()
	fmt.Println("main execution stopped")
}

// Run the program: go run main.go
 
//** EXPECTED OUTPUT: **//
// main execution started
// No. of Goroutines: 2
// f2 execution started
// f1(goroutine) execution started
// f1, i= 0
// f2(), i= 5
// f2(), i= 6
// f2(), i= 7
// f2 execution finished
// f1, i= 1
// f1, i= 2
// f1 execution finished
// main execution stopped