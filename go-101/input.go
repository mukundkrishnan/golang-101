package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	_, _ = fmt.Scan(&str)
	println(str)
	t, _ := strconv.ParseInt(str, 10, 0)
	println(t)
}
