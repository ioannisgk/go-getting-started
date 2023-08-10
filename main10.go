package main
import (
	"fmt"
)

type Circle struct {
    radius float64
    area   float64
}

type Student struct {
    name   string
    grades []int
}

func (s *Student) displayName() {
    fmt.Println(s.name)
}

func (s *Student) calcAverage() float64 {
    sum := 0
    for _, v := range s.grades {
        sum += v
    }
    return float64(sum * 100) / float64(len(s.grades) * 100)
}

func (c *Circle) calcArea1() {
    c.area = 3.14 * c.radius * c.radius
}

func (c Circle) calcArea2() {
    c.area = 3.14 * c.radius * c.radius
}

func main() {

    c1 := Circle{radius: 5}
    c1.calcArea1()
    fmt.Printf("%+v", c1)

    c2 := Circle{radius: 5}
    c2.calcArea2()
    fmt.Printf("%+v", c2)

    s := Student {
        name: "Joe",
        grades: []int{90, 75, 80},
    }
    s.displayName()
    fmt.Printf("%.2f%%", s.calcAverage())

}