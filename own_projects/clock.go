package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		hour, min, sec := time.Now().Clock()
		fmt.Printf("\rTime: %.2d:%.2d:%.2d", hour, min, sec)
	}
}
