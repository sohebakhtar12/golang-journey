package main

import "fmt"

func main() {

	// age := 20
	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// }else{
	// 	fmt.Println("person is not an adult")
	// }

	// age:=1
	// if age>=18 {
	//      fmt.Println("Person is an adult")
	// }else if age>=12{
	// 	fmt.Println("person is teenager")
	// }else{
	// 	fmt.Println("person is a kid")
	// }

	// role := "admin"
	// hasPermission := false
	// if role == "admin" && hasPermission {
	// 	fmt.Println("Yes")
	// }

	// we can declare a variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go does not have ternary, you will have to use normal if else

}
