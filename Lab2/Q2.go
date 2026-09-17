package main

import "fmt"

func main(){
	var n int 
	fmt.Print("Enter no. of students: ")
	fmt.Scan(&n)

	students := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter student %d name: ", i+1)
		fmt.Scan(&students[i])
	}

	fmt.Println("Initial Slice:", students)

	//add
	var name string
	fmt.Print("Enter student name to add: ")
	fmt.Scan(&name)

	students = append(students, name)
	fmt.Println("After Adding:", students)

	// Remove by index
	var index int
	fmt.Print("Enter index to remove: ")
	fmt.Scan(&index)

	if index >= 0 && index < len(students) {
		students = append(students[:index], students[index+1:]...)
		fmt.Println("After Removing:", students)
	} else {
		fmt.Println("Invalid index")
	}

	// Update
	fmt.Print("Enter index to update: ")
	fmt.Scan(&index)

	if index >= 0 && index < len(students) {
		fmt.Print("Enter new student name: ")
		fmt.Scan(&name)

		students[index] = name
		fmt.Println("After Updating:", students)
	} else {
		fmt.Println("Invalid index")
	}



	//part 2 map 
	marks := make(map[string]int)

	var subject string
	var mark int

	fmt.Print("\nEnter subject to insert: ")
	fmt.Scan(&subject)

	fmt.Print("Enter marks: ")
	fmt.Scan(&mark)

	marks[subject] = mark
	fmt.Println("After Inserting:", marks)

	// Delete
	fmt.Print("Enter subject to delete: ")
	fmt.Scan(&subject)

	delete(marks, subject)
	fmt.Println("After Deleting:", marks)

	// Lookup
	fmt.Print("Enter subject to lookup: ")
	fmt.Scan(&subject)

	mark, exists := marks[subject]

	if exists {
		fmt.Println("Marks:", mark)
	} else {
		fmt.Println("Subject not found")
	}


}



