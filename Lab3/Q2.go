package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func modifyValue(num *int) {
	*num = 20
}

func modifyStructure(s *Student){
	fmt.Println("Enter Name: ")
	fmt.Scan(&s.Name);

	fmt.Println("Enter Age: ")
	fmt.Scan(&s.Age);

}

func main() {

	// Part 1.
	x := 10

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value using pointer:", *(&x))

	// Part 2.
	fmt.Println("Value Before Modification call:", x)

	modifyValue(&x)

	fmt.Println("Value After Modification call:", x)

	// Part 3.
	student := new(Student)

	//fmt.Println(student) //call all the values present in student pointer object

	fmt.Println("Student details before Modification:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)

	modifyStructure(student)

	fmt.Println("Student details after modification: ")
	fmt.Println("Name: ", student.Name)
	fmt.Println("Name: ", student.Age)



	
	
}