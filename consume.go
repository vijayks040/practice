package main

import (
	"fmt"
	"time"
)

//var msgs = make(chan Person)

func consume() {
	if len(msgs) == 0 {
		fmt.Println("empty channel")
	}
	for {
		select {
		case msg := <-msgs:
			fmt.Println("deque ", "name:", msg.Name, "age: ", msg.Age, " timestamp:", msg.Timestamp)
		case <-time.After(5 * time.Second):
			done <- true
		}
	}
	//	for msg := range msgs {
	//		//msg := <-msgs
	//		fmt.Println("deque ", "name:", msg.Name, "age: ", msg.Age)
	//	}
	//wg.Done()
}
