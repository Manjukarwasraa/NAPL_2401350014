package main

import "fmt"

type Student struct{
	Name string
	Age int
	Marks float64
}

func main(){
	//create instance
	Manju := Student{
		Name : "Manju",
		Age : 19,
		Marks : 95.5,
	}
	fmt.Println(Manju)

	fmt.Printf("Name of Student: %s\n", Manju.Name)
	fmt.Printf("Age of Student: %d\n", Manju.Age)
	fmt.Printf("Marks of Student: %.2f\n", Manju.Marks)
}
