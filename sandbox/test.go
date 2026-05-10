package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) getName() {
	fmt.Println(p.Name)
}

func main() {
	p := &Person{
		Name: "cello",
	}
	p.getName()
}
