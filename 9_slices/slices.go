package main

import (
	"fmt"
	
)

//slice->dynamic array
//most used construct in go
//access useful methods

func main(){

	//uninitialized slice is nill

	// var arr []int          //initialize

	// fmt.Println(arr)       // op->[]
	// fmt.Println(arr==nil)  // op->true

	// fmt.Println(len(arr))
	                
	// var nums = make([]int,0 ,6)  //initialize another way 
	// fmt.Println(nums)            //op->[0 0 0 0 0] not nil [] (array)

	//capacity->maximum numbers of element can fit
	// fmt.Println(cap(nums))

	// element add
	// nums =append(nums, 1)  //add element in the last
	// nums =append(nums, 5)  //add element in the last
	// nums =append(nums, 9)  //add element in the last
	// nums =append(nums, 4)  //add element in the last
	// nums =append(nums, 6)  //add element in the last

	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//copy add to another array

	// var nums =make([]int,0,5)
	// nums=append(nums,3)
	// nums=append(nums,4)
	// var nums2 = make([]int,len(nums))

	// copy(nums2,nums)  //copy function
	
	// fmt.Println(nums)
	// fmt.Println(nums2)


	//slice operator

	// var nums = []int{1,2,3,4,5,6}
	// fmt.Println(nums[0:3])    // op->[1 2 3]

	//slice
	// var num1=[]int {1,2,3,4,5,6}
	// var num2=[]int {1,2,3,4,5,6}
    // fmt.Println(slices.Equal(num1,num2))  //true

	// 2D slice
	var nums = [][]int {{1,2,3}, {4,5,6}}
	fmt.Println(nums)



}