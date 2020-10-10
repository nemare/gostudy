package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Path []Point

type T struct {
	str string
}

func main() {
	var p = Point{1.2, 2.3}
	var q = Point{3.4, 11.2}
	fmt.Println(Distance(p, q))
	fmt.Println(q.Distance(p))

	t := T{"old Name"}
	t.getName()
	fmt.Println(t)
	t.getName1()
	fmt.Println(t)
	t1 := new(T)
	fmt.Println(t1)

	distanceFromQ := p.Distance
	fmt.Println(distanceFromQ(q))

	distanceFromQ1 := Point.Distance
	fmt.Println(distanceFromQ1(p, q))

	distanceFromQ2 := (*Point).Distance
	fmt.Println(distanceFromQ2(&p, q))

	path := Path{
		{1, 2},
		{4, 5},
	}

	anotherPoint := Point{5, 6}

	op := Path.translateBy

	//op(path, anotherPoint, true)

	op(path, anotherPoint, false)

}

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 实参
func (t T) getName() {
	t.str = "new Name"
}

// 形参
func (t *T) getName1() {
	t.str = "new Name"
}

func (p Point) Add(anotherPoint Point) Point {
	return Point{p.x + anotherPoint.x, p.y + anotherPoint.y}
}

func (p Point) Sub(anotherPoint Point) Point {
	return Point{p.x - anotherPoint.x, p.y - anotherPoint.y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.x, p.y)
}

func (path Path) translateBy(anotherPoint Point, add bool) {
	var op func(p, q Point) Point
	if add == true {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], anotherPoint)
		path[i].Print()
	}

}
