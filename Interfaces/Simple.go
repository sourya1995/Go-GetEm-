import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int){
	p.i = u
}

func fI(it Simpler) int {
	it.Set(5)
	return it.Get()
}

func main(){
	var s Simple
	fmt.Println(fI(&s))
}