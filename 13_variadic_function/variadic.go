package main

import "fmt"
//A variadic parameter allows a function to accept any number of arguments
func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
func main() {
	result := sum(5, 9, 6, 1, 3, 5, 8)
	fmt.Println(result)

}
