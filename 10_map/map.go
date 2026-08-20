package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map 
	// m := make(map[string]string)

    //setting an element
	// m["name"] = "soheb Akhtar"
    // m["area"]="backend"

    //get an element
	// fmt.Println(m["name"],m["area"])
	//Note->if key does not exists in the map then it return zero
	// fmt.Println(m["phone"])
	
	//find map length
	// fmt.Print(len(m))
    
	//delete key in the map
	// delete(m,"area")
	// fmt.Println(m)

	// delete map
	// clear(m)
	// fmt.Println(m)

	//create map in short
	// m:=map[string]int{"price": 40, "item": 5, "bill" : 500}
	// fmt.Println(m)

	// how to check element are present in the map
	r, ok :=m["price"]
	fmt.Println(r)  //it gives value and "ok" gives trues and false
	if ok {
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

	// m:=map[string]int{"price": 40, "item": 5, "bill" : 500}
    // m1:=map[string]int{"price": 40, "item": 5, "bill" : 500}
	// fmt.Println(maps.Equal(m,m1))  //true


	



}
