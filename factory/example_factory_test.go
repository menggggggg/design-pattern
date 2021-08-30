package factory

import "fmt"

func ExampleMakeProduct() {
	fmt.Println(MakeProduct("1").showName())
	fmt.Println(MakeProduct("2").showName())
	// Output:
	// Name is concreteProduct1
	// Name is concreteProduct2
}
