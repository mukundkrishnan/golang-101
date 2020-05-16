package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0, 3)
	for true {
		fmt.Print("Please enter a number(to Quit type X): ")
		var str string
		_, _ = fmt.Scan(&str)
		if strings.EqualFold(str, "x") {
			fmt.Println("You entered x. Goodbye!")
			break
		} else {
			num, err := strconv.Atoi(str)
			if err == nil {
				sli = append(sli, num)
				sort.Slice(sli, func(i, j int) bool { return sli[i] < sli[j] })
			}
		}
		fmt.Print("Sorted: ")
		for _, k := range sli {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}
}
