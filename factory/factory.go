package factory

import (
	"fmt"
)

type Product interface {
	show() string
}

type ConcreteProduct1 struct {
	name string
}

func (c *ConcreteProduct1) show() string {
	fmt.Printf("ConcreteProduct1 Name is %s \n", c.name)
	return c.name
}

type ConcreteProduct2 struct {
	name string
}

func (c *ConcreteProduct2) show() string {
	fmt.Printf("ConcreteProduct2 Name is %s \n", c.name)
	return c.name
}

func MakeProduct(kind string) Product {
	var p Product
	switch kind {
	case "1":
		p = &ConcreteProduct1{name: "concreteProduct1"}
	case "2":
		p = &ConcreteProduct2{name: "concreteProduct2"}
	default:
		p = nil
	}
	return p
}
