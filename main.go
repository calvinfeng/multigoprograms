package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")
	for range time.NewTicker(time.Second).C {
		fmt.Println("current time is", time.Now())
	}
}
