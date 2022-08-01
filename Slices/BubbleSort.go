package main
import "fmt"

func main() {
	sla := []int{2,6,4,-10,8,89,12,68,-45,37}
	fmt.Println("before sort: ", sla)
	bubbleSort(sla)
	fmt.Println("after sort: ", sla)
}

func bubbleSort(sl []int){
	for pass := 1; pass < len(sl); pass++ {
		for i := 0; i < len(sl) - pass; i++ {
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
}