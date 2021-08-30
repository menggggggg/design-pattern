package factory

import (
	"fmt"
)

type Product interface {
	showName() string
}

type ConcreteProduct1 struct {
	name string
}

func (c *ConcreteProduct1) showName() string {
	return fmt.Sprint("Name is ", c.name)
}

type ConcreteProduct2 struct {
	name string
}

func (c *ConcreteProduct2) showName() string {
	return fmt.Sprint("Name is ", c.name)
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
