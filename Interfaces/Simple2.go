package main
import (
	"fmt"

)

type Simpler interface{
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int{
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

type RSimple struct {
	i int
	j int
}

func (p *RSimple) Get() int {
	return p.j
}

func(p *RSimple) Set(u int) {
	p.j = u
}

func fI(it Simpler) int {
	switch it.(type){
	case *Simple:
		it.Set(5)
		return it.Get()
	

	case *RSimple:
		it.Set(50)
		return it.Get()

	default:
		return 99
	}
	return 0

}

func main(){
	var s Simple
	fmt.Println(fI(&s))
	var r RSimple
	fmt.Println(fI(&r))
}