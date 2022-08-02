package main
import (
	"container/list"
	"fmt"
)

func insertListElements(n int)(*list.List){
	lst := list.New()
	for i:=1; i<=n; i++ {
		lst.Pushback(i)
	}
	return lst
}

func main(){
	n := 5
	myList := insertListElements(n)
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}