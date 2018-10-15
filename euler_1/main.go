package main

import "fmt"

func main() {
	fmt.Printf("sum -> %v", getMultiple(3, 5,1000))
}

func getMultiple(x int, y int,limit int) int {
	sum := 0
	mp := 1
	m := mp * x
	for m < limit {
		sum += m
		mp += 1
		m = mp * x
	}
	mp = 1

	m = mp * y
	for m < limit {
		if m % x !=0{
			sum += m
		}
		mp += 1
		m = mp * y
	}

	return sum
}
