package main

import "fmt"

//Basic Funcrion Syntax
// func functionName(parameters) returnType {
//     // code
// }

//(a int, b int) both are same
// func add(a , b int) int {
// 	return a + b
// }

// func getLuanguage()(string,string,string){
// 	return "golang","javastript","java"
// }

//reverse array
func reverseArr(arr[]int)[]int{
	left := 0
	right :=len(arr)-1
	for left<=right {
		arr[left],arr[right]=arr[right],arr[left]
		left ++
		right--

	}
	return arr
}

func main() {
	// result := add(9, 12)
	// fmt.Println(result)
   
	// fmt.Println(getLuanguage())

	// lang1, lang2,_:=getLuanguage()
	// fmt.Println(lang1,lang2)  //op->golang javastript

	//reverse array
	arr := []int{1,2,3,4,5,6}
	fmt.Println(reverseArr(arr))
}
