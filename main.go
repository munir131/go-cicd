package main

import "fmt"

func main() {
	var n1 int32 = 4
	var n2 int32 = 5
	total := Add(n1, n2)
	fmt.Println(total)
}

func Add(num1, num2 int32) int32 {
	return num1 + num2
}
