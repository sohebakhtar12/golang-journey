package main

import "fmt"

func main() {
	arr := []int{5, 9, 6, 3, 8, 2}
    //using for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//using range
	// sum:=0
    // for i,num:=range arr{  //i->index and num->element
		// fmt.Println(num)   //print element
		// fmt.Println(i)   //print index
		// sum+=num
	// }
	// fmt.Println(sum)      //print sum of element

	m:=map[string]int {"item" : 5 ,"price":50,"discount": 5,"saleItem":10}

	for key,val :=range m{
		fmt.Println(key,val)
	}


}