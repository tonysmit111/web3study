package main

import (
	"fmt"
	"math"
)

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Long float64
	Wide float64
}

func (r Rectangle) Area() float64 {
	return r.Long * r.Wide
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Long + r.Wide)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var r Shape
	r = Rectangle{
		Long: 2,
		Wide: 3,
	}
	fmt.Println("面积：", r.Area())
	fmt.Println("周长", r.Perimeter())

	c := Circle{Radius: 3}
	var cs Shape = &c
	fmt.Println("面积：", cs.Area())
	fmt.Println("周长", cs.Perimeter())

	var cs2 Shape = &Circle{Radius: 1}
	fmt.Println("面积：", cs2.Area())
	fmt.Println("周长", cs2.Perimeter())
}
