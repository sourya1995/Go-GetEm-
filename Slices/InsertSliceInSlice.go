package main
import "fmt"

func main{
	s := []string {"M", "N", "O", "P", "Q", "R"}
	in := []string {"A", "B", "C"}
	res := insertSlice(s, in, 0) //front insertion
	fmt.Println(res)
	res := insertSlice(s, in, 3) //middle insertion
	fmt.Println(res)
}

func insertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice) + len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

/* slice - slice in which another slice is inserted
	insertion - slice that is to be inserted
	index - point of insertion
	return the updated slice

	create a slice result, using make function
	length of result = length of "slice" + length of "insertion"
	inserting at "index" means, elements of slice starting at that index + after that index, will move
	len(insertion) forward

	first, copy all contents from 0, till index, that is, the element just before index, to result.
	then, copy "insertion" into result, AFTER the contents from slice
	now, we copy remaining elements from slice, if any, to result
	so - copy(result[at:], slice[index:]) - everything after "at" in result, with everything after index: in 
	"slice"
	return result