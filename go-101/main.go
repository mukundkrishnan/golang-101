package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a String: ")
	if scanner.Scan() {
		fmt.Printf("You wrote \"%s\"\n", scanner.Text())
	}

	var str = scanner.Text()
	if strings.HasPrefix(strings.ToLower(str), "i") &&
		strings.HasSuffix(strings.ToLower(str), "n") &&
		strings.Contains(strings.ToLower(str), "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
