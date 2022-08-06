package main
import ("fmt")

func main() {
	ch := make(chan int)

	go func() {
		for {
			fmt.Print(<-ch, " ")
		}
	}()

	for i := 0; i<=100000; i++ {
		select{
		case ch <- 0:
		case ch <- 1:
		}
	}
}