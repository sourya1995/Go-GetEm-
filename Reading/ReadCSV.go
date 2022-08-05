/* delimiter is ; */

package main
import (
	"bufio"
	"fmt"
	"log"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title string
	price float64
	quantity int
}

func main() {
	bks := make([]Book, 1) //make a slice with a length of 1
	file, err := os.Open("products.txt") //open products.txt
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close() //close the file after open

	reader := bufio.NewReader(file) //new reader object on file
	for { //infinite for loop
		line, err := reader.ReadString('\n') //read a line
		if err == io.EOF{
			break
		} //jump out of for loop when EOF is reached
		line = string(line[:len(line) - 2]) //len(line) - 2 because \n is also a character in string
		strS1 := strings.Split(line, ";") //read the line, split on ;, store in array StrS1
		book := new(Book)
		book.title = strS1[0] //first element is the title
		book.price, err = strconv.ParseFloat(strS1[1], 32) //price is second field- we convert string value to float
		if err != nil {
			fmt.Printf("Error in file: %v", err)

		}
		book.quantity, err = strconv.Atoi(strS1[2]) //quantity is integer, so we have to convert string to integer
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		if bks[0].title == "" { 
			bks[0] = *book
		} else{
			bks =  append(bks, *book) //if there is a title, append to slice
		}
	}

	fmt.Println("we have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}