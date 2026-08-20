package main

import "fmt"
//for -> only construct in go for looping
func main() {

	//while loop using for 	
	i:= 1
	for  i <= 3 {

		fmt.Println(i)
		i++
	}

//infinite loop
// for{
// 	fmt.Println(1)
// }	

//classic for loop

// for i:=0; i<3 ;i++{
// 	fmt.Println(i)
// }

//1.22 range 

// for i := range 5 {
// 	fmt.Println(i)
// }

}