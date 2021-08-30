package abstract_factory

type FacoryProducer struct {
}

func (fp *FacoryProducer) getFacory(typ string) AbstractFactory {
	if typ == "shape" {
		return &ShapeFactory{}
	} else if typ == "color" {
		return &ColorFactory{}
	} else {
		return nil
	}
}

type AbstractFactory interface {
	getShape(string) Shape
	getColor(string) Color
}

type ShapeFactory struct {
}

func (sf *ShapeFactory) getShape(shape string) Shape {
	if shape == "circle" {
		return &Circle{}
	} else if shape == "square" {
		return &Square{}
	} else if shape == "rectangle" {
		return &Rectangle{}
	} else {
		return nil
	}
}

func (sf *ShapeFactory) getColor(color string) Color {
	return nil
}

type ColorFactory struct {
}

func (cf *ColorFactory) getShape(shape string) Shape {
	return nil
}

func (cf *ColorFactory) getColor(color string) Color {
	if color == "red" {
		return &Red{}
	} else if color == "green" {
		return &Green{}
	} else if color == "blue" {
		return &Blue{}
	} else {
		return nil
	}
}

type Shape interface {
	draw() string
}

type Circle struct {
}

func (c *Circle) draw() string {
	return "draw circle!"
}

type Square struct {
}

func (s *Square) draw() string {
	return "draw square!"
}

type Rectangle struct {
}

func (r *Rectangle) draw() string {
	return "draw rectangle!"
}

type Color interface {
	fill() string
}

type Red struct {
}

func (r *Red) fill() string {
	return "fill red!"
}

type Green struct {
}

func (g *Green) fill() string {
	return "fill green!"
}

type Blue struct {
}

func (b *Blue) fill() string {
	return "fill Blue!"
}
