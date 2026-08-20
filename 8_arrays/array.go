package main

import "fmt"

//number sequence of specific length
func main(){
	var arr[5]int
	fmt.Println(arr)   //[0 0 0 0 0] zero values

	arr[0]=2
	fmt.Println(arr)   //[2 0 0 0 0] zero values

	//find array length
	n:=len(arr)
	fmt.Println(n)

	var vals [4]bool
	fmt.Println(vals)   //[false false false false]

    vals[2]=true
	fmt.Println(vals)   //[false false true false]

	var name [5]string 
	name[0] = "soheb"
	fmt.Println(name)

		// to declare it in single line
	// nums := [3]int{1, 2,7}
	// fmt.Println(nums)
	// fmt.Println(nums[1])   //access array element

	// 2d arrays
	nums := [2][2]int{{3, 4}, {5,7}}
	fmt.Println(nums)

	// - fixed size, that is predictable
	// - Memory optimazation
	// - Contant time access


}