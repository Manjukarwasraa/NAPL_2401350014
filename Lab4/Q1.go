package main

import "fmt"

func square(n int, ch chan string){
	fmt.Println("Square goroutine start")
	ch <- fmt.Sprintf("Square = %d", n*n)
	fmt.Println("Square goroutine completed")
}


func cube(n int, ch chan string){
	fmt.Println("Square goroutine start")
	ch <- fmt.Sprintf("Cube = %d", n*n*n)
	fmt.Println("Cube goroutine completed")
}

func fibonacci(n int, ch chan string) {
    fmt.Println("Fibonacci goroutine started")

    a, b := 0, 1
    for i := 0; i < n; i++ {
        a, b = b, a+b
    }

    ch <- fmt.Sprintf("Fibonacci number = %d", a)
    fmt.Println("Fibonacci goroutine completed")
}

func main() {
    var n int

    fmt.Print("Enter a number: ")
    fmt.Scan(&n)

    ch := make(chan string)

    go square(n, ch)
    go cube(n, ch)
    go fibonacci(n, ch)

    for i := 0; i < 3; i++ {
        fmt.Println(<-ch)
    }

    fmt.Println("Main completed")
}