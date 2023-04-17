package main

import (
	"fmt"
	"math"
)

type areaShape interface {
	area() float64
}

type perimShape interface {
	perim() float64
}

type circle struct {
	radius float64
}

type tria struct {
	sideA  float64
	sideC  float64
	base   float64
	height float64
}

type trape struct {
	base1  float64
	base2  float64
	sideC  float64
	sideD  float64
	height float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (t *trape) perim() float64 {
	return t.base1 + t.base2 + t.sideC + t.sideD
}

func (tr *tria) perim() float64 {
	return tr.sideA + tr.sideC + tr.base
}

func (tr *tria) area() float64 {
	return tr.base * tr.height / 2
}

func (t *trape) area() float64 {
	return (t.base1 + t.base2) * t.height / 2
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	c1 := circle{3.5}
	r1 := rect{4, 5}
	t1 := trape{4, 5, 6, 7, 8}
	tr1 := tria{3, 5, 8, 6}

	areaShapes := []areaShape{&c1, &r1, &t1, &tr1}
	perimShapes := []perimShape{&t1, &r1, &tr1}

	fmt.Println("Area:")

	for _, areaShape := range areaShapes {
		fmt.Println(areaShape.area())
	}

	fmt.Println("----------")

	fmt.Println("Perimeter:")

	for _, perimShape := range perimShapes {
		fmt.Println(perimShape.perim())
	}
}
