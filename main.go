package main

import (
	"fmt"
	"github.com/epond/go-sketchpad/string"
	"time"
)

func main() {
	fmt.Printf("Hello, %v\n", string.Reverse("World"))

	i := 0
	c := time.Tick(time.Duration(1) * time.Second)
	fmt.Println("Doing work")
	for _ = range c {
		i += 1
		// should we do work?
		fmt.Println("Doing work")
		// should we break from the loop?
		if i > 5 {
			break
		}
	}

}
