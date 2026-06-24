package main

import (
	"fmt"
	"sync"
	"time"
)

type PostCounter struct {
	counter int
	sync    sync.Mutex
}

func (pc *PostCounter) Increment() {
	pc.sync.Lock()
	defer pc.sync.Unlock()
	pc.counter++
}

func (pc *PostCounter) GetCount() int {
	return pc.counter
}

func SecureCounter() {
	newPostCounter := PostCounter{}

	for i := 1; i <=1000; i++ {
		go func(){
			newPostCounter.Increment()
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Count is ", newPostCounter.counter)
}


/*
Scenario: You are building a hit counter for a native Nigerian blogs platform (like a tech version of Linda Ikeji). When an article blows up, thousands of users hit the page at once.
Your Task: Create a struct called PostCounter that holds an integer count and a sync.Mutex. Write two methods on this struct using pointer receivers: Increment() to safely add 1 to the count, and GetCount() int to safely read the current total.
Goal: Understand basic struct attachment, pointers, and the standard Lock() / Unlock() flow.
*/
