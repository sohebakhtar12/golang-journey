package main

import "fmt"

func main() {

	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println(("three"))
	// default:
	// 	fmt.Println("other")
	// }

	//Multiple condition switch
	// switch time.Now().Weekday(){
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("it's weekday")
	// default:
	// 	fmt.Println("it's workday")
	// }

	//type switch
	whoeAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its an string")
		case bool:
			fmt.Println("its an bool")
		case float32:
			fmt.Println("its an float")
		default:
			fmt.Println("other")			
		}
	}
	whoeAmI(true)

}
