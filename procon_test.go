package main

import (
	"testing"
)

// A test is created by writing a function with a name
// beginning with `Test`.
func Testprocon(t *testing.T) {
	go produce(3)
	go consume()
}
