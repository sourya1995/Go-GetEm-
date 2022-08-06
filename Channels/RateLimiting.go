package main

const (
	AvailableMemory = 10 << 20
	AverageMemoryPerRequest = 10 << 10
	MAXREQS = AvailableMemory / AverageMemoryPerRequest	
)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b int
	replyc chan int
}

func process(r *Request){

}

func handle(r *Request){
	process(r)
	<- sem
}

func Server(queue chan *Request){
	for{
		sem <- 1

		request := <- queue
		go handle(request)
	}
}

func main() {
	queue := make(chan *Request)
	go Server(queue)
}