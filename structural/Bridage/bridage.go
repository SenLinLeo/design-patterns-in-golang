package Bridage

import "fmt"

type Draw interface {
	DrawCircle(radius, x, y int)
}

type RedCircle struct {
}

func (r *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("radius x y ", radius, x, y)
}

type YellowCircle struct {
}

func (r *YellowCircle) DrawCircle(radius, x, y int) {
	fmt.Println("raduis x y", radius, x, y)
}

type Shape struct {
	draw Draw
}

func (s *Shape) Shape(d Draw) {
	s.draw = d
}

type Circle struct {
	shape Shape
	x int
	y int
	raduis int
}

func (c *Circle) Constructor(x, y, radius int, draw Draw) {
	c.x = x
	c.y = y
	c.raduis = radius
	c.shape.Shape(draw)
}

func (c *Circle) Draw() {
	c.shape.draw.DrawCircle(c.raduis, c.x, c.y)
}