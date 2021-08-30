package abstract_factory

import "fmt"

func ExampleAbstractFactory() {
	producer := &FacoryProducer{}
	abstractFactory := producer.getFacory("shape")
	circle := abstractFactory.getShape("circle")
	fmt.Println(circle.draw())

	abstractFactory = producer.getFacory("color")
	red := abstractFactory.getColor("red")
	fmt.Println(red.fill())
	// Output:
	// draw circle!
	// fill red!
}
