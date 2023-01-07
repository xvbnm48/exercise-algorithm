package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func (p *Person) Birthday() {
	p.age++
}

func (p Person) Greet() string {
	return "Hello guys, my name is " + p.name + " and I am " + strconv.Itoa(p.age) + " years old."
}

func main() {
	p := Person{"Sakura Miyawaki", 20}
	p.Birthday()
	fmt.Println(p.Greet())
}
