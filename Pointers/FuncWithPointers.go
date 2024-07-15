package main

import (
	"fmt"
)

func add5Value(count int){
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5Point(count *int){
	*count += 5
	fmt.Println("add5Value: ", *count)
}

func main(){
	var count int
	add5Value(count)
	fmt.Println(count)
	add5Point(&count)
	fmt.Println(count)
}