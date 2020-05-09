package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int64, 0, 3)
	var num int64
	for true {
		fmt.Print("Please enter a number(to Quit type x): ")
		var str string
		_, _ = fmt.Scanf("%s", &str)
		if strings.EqualFold(str, "x") {
			fmt.Println("You entered x. Goodbye!")
			break
		} else {
			num, _ = strconv.ParseInt(str, 10, 0)
			sli = append(sli, num)
			sort.Slice(sli, func(i, j int) bool { return sli[i] < sli[j] })
		}
		fmt.Print("Sorted: ")
		for _, k := range sli {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}
}
