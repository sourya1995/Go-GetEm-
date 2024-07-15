package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Println("at least %v arguments are needed", minArgs)
		os.Exit(1)
	}

	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func main(){
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("the longest word passed was:", longest)

	}
	else {
		fmt.Println("There was an error")
		os.Exit(1)
	}
}