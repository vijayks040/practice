package main

import (
	"fmt"
	"sync"
	"time"
)

//struct to be compatible with json marshall and unmarshalling
type Person struct {
	Name      string    `json:"name,omitempty"`
	Age       int       `json:"category1,omitempty"`
	Timestamp time.Time `json:"time,omitempty"`
}

//channels to receive values sent from producer
var msgs = make(chan Person)

var done = make(chan bool) //this bool channel is used to signal that the job is finished
//var wg sync.WaitGroup//tried it before implementing channel for signalling finish

func main() {
	fmt.Println("starting now...")
	//wg.Add(2)
	go produce() //calling produce function concurrently with other functions
	go consume()
	//wg.Wait()
	time.Sleep(2 * time.Second) //to check in between writing
	go produce_duplicate()
	<-done
	fmt.Println("exiting gracefully...")
}
