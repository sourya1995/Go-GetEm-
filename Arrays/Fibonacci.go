package main
import "fmt"

var fib[10]int64

func fibs() [10]int64{
	fib[0] = 1
	fib[1] = 1

	for i := 2; i<10; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func main(){

	arr := fibs()
	for i := 0; i < 10; i++ {
		fmt.Printf("The %d-th Fibonacci number is: %d", i, arr[i])
	}
}