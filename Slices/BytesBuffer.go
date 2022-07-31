package main
import (
	"fmt"
	"bytes")

func main(){

	var b bytes.Buffer

	b.WriteString("ABC")
	b.WriteString("DEF")

	fmt.Fprintf(&b, "A number: %d, a string: %v", 10, "bird")
	b.WriteString("[DONE]")
	fmt.Println(b.String())
}

