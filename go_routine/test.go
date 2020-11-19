package main

import (
	"fmt"
	"time"
)

func addNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func addChar() {
	for c := 'a'; c <= 'e'; c++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%c ", c)
	}
}

func main() {
	fmt.Println("test begin")

	go addNumber()
	go addChar()

	time.Sleep(1 * time.Second)

	fmt.Println("test end")
}
