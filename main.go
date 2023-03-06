package main

import "fmt"

func main() {
	var n1 int32 = 4
	var n2 int32 = 5
	total := Add(n1, n2)
	fmt.Println(total)

	var n3 float32 = 4.5
	var n4 float32 = 5.5
	totalFloat := Add(n3, n4)
	fmt.Println(totalFloat)
}

func Add[N int32 | float32](num1, num2 N) N {
	return num1 + num2
}
