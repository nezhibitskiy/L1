package main

import "fmt"

func main() {
	var data int64

	data = data | 1<<3
	fmt.Println(data)
}
