package main

import "fmt"

const Pi = 3.14159265358

func main() {
	fmt.Println("Hello World")
	var intNum int
	fmt.Println(intNum)


	var list []int = make([]int, 10)
	fmt.Println(list)
	fmt.Println(splice(list))

	var radius float64 = 2
	fmt.Println(findArea(radius))

	fmt.Println(multiples(1, 2,3,4,5))
}

func splice(list []int) []int {
	for i := range len(list) {
		list[i]++;
	}
	return list
}

func findArea(radius float64) float64 {
	return radius*radius*Pi
}

func multiples(nums ...int) int {
	fmt.Println("original splice", nums, " ")
	total := 1

	for _, val := range nums {
		total *= val
	}
	return total
}