package main

import "fmt"

func main() {
	var (
		ab    uint = 100
		ac    uint = 200
		count uint = 10
	)
	for count > 0 {
		if ab != ac {
			fmt.Println("cool")
		} else if ab < ac {
			fmt.Println("down")
		}
		count--
	}
}

