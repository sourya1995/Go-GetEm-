package main
import "fmt"

func main(){
	value := 0
	screen := [2][2]int{}

	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = value
			value = value + 1
		}
	} //assigning values

	for row := range screen {
		for column := range screen[0] {
			fmt.Print(screen[row][column], " ")
		}
		fmt.Print("\n")
	}
}