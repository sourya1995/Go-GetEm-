package main

import(
	"fmt"
	"./mysort"
)

type Person struct{
	firstName string
	lastName string
}

type Persons []Person

func (p Persons) Len() int {return len(p)}

func(p Persons) Less(i, j int) bool {
	in := p[i].lastName + " " + p[i].firstName
	jn := p[j].lastName + " " + p[j].firstName
	return in < jn
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}