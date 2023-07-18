package main

import "fmt"

func superAdd(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// 자바 같은 방식의 for loop
func forLoop(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	sum := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)
	sum_1 := forLoop(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum_1)
}
