package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float32
  perimeter() float32
}

type Circle struct {
  radius float32
}

type Rectangle struct {
  w, h float32
}

type Square struct {
  a float32
}

type Triangle struct {
  w, h float32
}

func (circ Circle) area() float32 {
  return math.Pi * circ.radius * circ.radius
}

func (rect Rectangle) area() float32 {
  return rect.h * rect.w
}

func (sqr Square) area() float32 {
  return sqr.a * sqr.a
}

func (tri Triangle) area() float32 {
  return tri.h * tri.w / 2
}

func (circ Circle) perimeter() float32 {
  return 2 * math.Pi * circ.radius
}

func (rect Rectangle) perimeter() float32 {
  return (rect.h * 2) +  (rect.w * 2)
}

func (sqr Square) perimeter() float32 {
  return sqr.a * 4
}

func (tri Triangle) perimeter() float32 {
  return 0
}

func getArea (shape Shape) float32 {
  return shape.area()
}

func getPerimeter (shape Shape) float32 {
  return shape.perimeter()
}

func main() {
  circle := Circle{radius: 5}
  rect := Rectangle{h: 5, w: 2}
  sqr := Square{a: 5}
  tri := Triangle{h: 5, w:2}

  fmt.Println("Circle area: ", getArea(circle), " Perimeter: ", getPerimeter(circle))
  fmt.Println("Rectangle area: ", getArea(rect), " Perimeter: ", getPerimeter(rect))
  fmt.Println("Square area: ", getArea(sqr), " Perimeter: ", getPerimeter(sqr))
  fmt.Println("Triangle area: ", getArea(tri))
}
