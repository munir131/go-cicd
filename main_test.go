package main

import "testing"

func TestAdd(t *testing.T) {
	var n1 int32 = 1
	var n2 int32 = 2
	var expected int32 = 3
	total := Add(n1, n2)
	if expected != total {
		t.Errorf("Ans should be %d", expected)
	}
}
