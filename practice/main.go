 package main

import (
	"fmt"
)


// func main() {
// 	fmt.Println("Hello,  Manju")

	
// }

// 	// var name string
// 	// var rollno int
// 	// name = "Manju"
// 	// rollno = 14

// 	// fmt.Printf(" Name: %s", name)

// 	// fmt.Println()
// 	// fmt.Println(" Rollno: %d ", rollno)

// 	// var a int = 66
// 	// b := string(a)

// 	// //string conversion give ascii value
// 	// fmt.Println(b)

// 	// var a int = 66
// 	// b := strconv.Itoa(a)
// 	// fmt.Println(b)

// 	//if we want the same int value in string so use strconv packAGE

// 	// var a string
// 	// a = "hello"

// 	// fmt.Println(len(a))

// 	// fmt.Println(len([]rune(a)))

// 	//for every char of string

// 	// for i, ch:= range "He`llo"{
// 	// 	if i == 0{
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println(i, string(ch))
// 	// }

// 	// var marks int = 60
// 	// switch {
// 	// 	case marks >= 60:
// 	// 		fmt.Println("First")

// 	//     case marks >= 50:
// 	// 	    fmt.Println("Second")

// 	//     case marks >= 40:
// 	// 	    fmt.Println("Third")

// 	//     default:
// 	// 	    fmt.Println("Fail")

// // 	var marks int

// // 	fmt.Print("Enter your marks: ")
// // 	fmt.Scan(&marks)

// // 	switch {
// // 	case marks >= 60:
// // 		fmt.Println("First")

// // 	case marks >= 50:
// // 		fmt.Println("Second")

// // 	case marks >= 40:
// // 		fmt.Println("Third")

// // 	default:
// // 		fmt.Println("Fail")
// // 	}


// }

// // go run filename.go



// package main

// import "fmt"

// func sum(nums ...int) int {
// 	total := 0

// 	for _, n := range nums {
// 		total += n
// 	}

// 	return total
// }

// func main() {
// 	fmt.Println("Hello, Manju")

// 	result := sum(10, 20, 30, 40)
// 	fmt.Println(result)

// package main

// import "fmt"


// type Celsius float64

// func (c Celsius) ToF() float64{
// 	return float64(c)*1.8 + 32
// }

// func main(){
// 	var temp Celsius = 37.4
// 	fmt.Println(temp.ToF())
// }


// package main

// import "fmt"

// func result(a, b, c int) (float64, bool) {
// 	avg := float64(a+b+c) / 3
// 	return avg, avg >= 40
// }

// func main() {
// 	avg, passed := result(10, 60, 80)

// 	fmt.Println("Average:", avg)
// 	fmt.Println("Passed:", passed)
// }