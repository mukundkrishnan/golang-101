package main

import "fmt"

func main() {
	x := [...]int {7,28,35,14}
	for i, v:= range x {
		fmt.Printf("ind %d, val %d\n", i, v)
	}
}
