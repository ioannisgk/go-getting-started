package main
import (
	"fmt"
)

type Student struct {
	name string
	rollNo int
}

type Circle struct {
	x int
	y int
	radius float64
	area float64
}

type struct1 struct {
	x int
}

type struct2 struct {
	x int
}

func calcArea1(c Circle) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	c.area = area
}

func calcArea2(c *Circle) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	(*c).area = area
}

func main() {

	var s1 Student
	fmt.Printf("%+v \n", s1)

	s2 := new(Student)
	fmt.Printf("%+v \n", *s2)

	s3 := Student {
		name: "Joe",
		rollNo: 12,
	}
	fmt.Printf("%+v \n", s3)

	s4 := Student{"Alex", 13}
	fmt.Printf("%+v \n", s4)

	c := Circle{
		x: 5,
		y: 5,
		radius: 5,
		area: 0,
	}
	fmt.Printf("%+v \n", c)
	calcArea1(c)
	fmt.Printf("%+v \n", c)
	fmt.Printf("%+v \n", c.area)

	calcArea2(&c)
	fmt.Printf("%+v \n", c)
	fmt.Printf("%+v \n", c.area)

	c1 := struct1{x: 5}
	c2 := struct1{x: 6}

	if c1 == c2 {
		fmt.Println("Same struct types, same values")
	} else {
	    fmt.Println("Same struct types, different values")
	}

}