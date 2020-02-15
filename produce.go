package main

import "time"

//this function takes care of sending values to the channel
func produce() {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second) //making sure each entry is separated
		msgs <- Person{
			Name:      "vijay",
			Timestamp: time.Now(),
			Age:       i}
		//msgs <- p1

	}
	//	p2 := Person{
	//		Name:    "vinay",
	//		Address: "channagiri",
	//		Age:     30}
	//	msgs <- p2
	//done <- true
}

//checking whether concurrent writing is working fine or not
func produce_duplicate() {
	msgs <- Person{
		Name:      "vvvvvvvvvvvvv",
		Timestamp: time.Now(),
		Age:       28}

}
