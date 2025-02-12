package main

import (
	"fmt"
	"math"
)

type Vec2i struct {
	X, Y int
}

func NewVec2i(x, y int) Vec2i {
	return Vec2i{x, y}
}

func (v Vec2i) Add(other Vec2i) Vec2i {
	return Vec2i{v.X + other.X, v.Y + other.Y}
}

func (v Vec2i) Sub(other Vec2i) Vec2i {
	return Vec2i{v.X - other.X, v.Y - other.Y}
}

func (v Vec2i) Mul(other Vec2i) Vec2i {
	return Vec2i{v.X * other.X, v.Y * other.Y}
}

func (v Vec2i) Norm() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vec2i) Normalized() Vec2i {
	norm := v.Norm()
	return Vec2i{int(float64(v.X) / norm), int(float64(v.Y) / norm)}
}

// Produit scalaire
func (v Vec2i) Dot(other Vec2i) int {
	return v.X*other.X + v.Y*other.Y
}

// Produit vectoriel
func (v Vec2i) Cross(other Vec2i) int {
	return v.X*other.Y - v.Y*other.X
}

func main() {
	// Example
	v1 := NewVec2i(1, 2)
	v2 := NewVec2i(3, 4)

	v3 := v1.Add(v2)
	fmt.Printf("v1 + v2 = %v\n", v3)

	v4 := v1.Sub(v2)
	fmt.Printf("v1 - v2 = %v\n", v4)
	v5 := v1.Mul(v2)
	fmt.Printf("v1 * v2 = %v\n", v5)

	norm := v1.Norm()
	fmt.Printf("Norm of v1 = %v\n", norm)

	normalized := v1.Normalized()
	fmt.Printf("Normalized v1 = %v\n", normalized)

	dot := v1.Dot(v2)
	fmt.Printf("v1 . v2 = %v\n", dot)

	cross := v1.Cross(v2)
	fmt.Printf("v1 x v2 = %v\n", cross)
}
