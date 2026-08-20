package main

import "fmt"

//slice->dynamic array
//most used construct in go
//access useful methods
func main(){
	//uninitialized slice is nill

	var arr []int

	fmt.Println(arr)      //[]
	fmt.Println(arr==nil)   //true

	fmt.Println(len(arr))
	                
	var nums = make([]int,0 ,6)  //[0 0 0 0 0] not nil [] (array)
	fmt.Println(nums)

	//capacity->maximum numbers of element can fit
	fmt.Println(cap(nums))

	// element add
	nums =append(nums, 1)  //add element in the last
	nums =append(nums, 5)  //add element in the last
	nums =append(nums, 9)  //add element in the last
	nums =append(nums, 4)  //add element in the last
	nums =append(nums, 6)  //add element in the last

	fmt.Println(nums)
	fmt.Println(cap(nums))
}