package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("Hello Worldj")
}

func main() {
	go sayHello()
	runtime.Gosched()
	fmt.Println("Main goroutine")
}
