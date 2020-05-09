package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the file name with Path(Example: /Users/mukund/test.txt): ")
	scanner := bufio.NewScanner(os.Stdin)
	var fileName string
	if scanner.Scan() {
		fileName = scanner.Text()
	}

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	sli := make([]Name, 0, 4)
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) > 1 {
			n := Name{fname: words[0], lname: words[1]}
			sli = append(sli, n)
		}
	}

	for _, s := range sli {
		fmt.Printf("FirstName: %s, LastName: %s\n", s.fname, s.lname)
	}

}
