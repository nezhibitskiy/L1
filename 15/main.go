package main

import (
	"fmt"
	"reflect"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if reflect.TypeOf(v) == reflect.TypeOf(justString) {
		runes := []rune(fmt.Sprintf("%v", v))
		if len(runes) < 100 {
			justString = fmt.Sprintf("%v", v)
			return
		}
		justString = string([]rune(fmt.Sprintf("%v", v))[:100])
	}
}

func createHugeString(i int) interface{} {
	return "erewgwg"
}

func main() {
	someFunc()
	fmt.Println(justString)
}
