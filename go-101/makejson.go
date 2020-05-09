package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a Name: ")
	if scanner.Scan() {
		m["name"] = scanner.Text()
	}
	fmt.Print("Enter an Address: ")
	if scanner.Scan() {
		m["address"] = scanner.Text()
	}
	jsonString, _ := json.Marshal(m)
	fmt.Println("Marshalled Json: ", jsonString)

	jsonData := string(jsonString)
	fmt.Println("Json Object: ", jsonData)
}
