package main
import (
	"fmt"
	"strings"
	"strconv"
)

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func returnCube(n int) int {
    return n*n*n
}

func operation1(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func operation2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
        sum += value
    }
	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey ", student, ", here are your subjects - ")
	for _, sub := range subjects {
        fmt.Printf("%s, ", sub)
    }
}

func f() (int, int) {
	return 42, 53
}

func calcSquare(numbers []int) []int {
    squares := []int{}
    for _, v := range numbers {
        squares = append(squares, v*v)
    }
    return squares
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}

func print(n int) {
    if n == 0 {
        return
    }
    print(n - 1)
    fmt.Print(n)

}

func main() {

	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)
    
	fmt.Println(returnCube(5))

	fmt.Println(operation1(20, 10))

	sum, diff := operation2(20, 10)
	fmt.Println(sum, " ", diff)

	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(100))
	fmt.Println(sumNumbers(100, 200, 300))

	printDetails("Joe", "Physics", "Biology")

	v, _ := f()
	fmt.Println(v)

	nums := [3]int{10, 20, 15}
    fmt.Println(calcSquare(nums[:]))

	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is:", result)

	print(5)

	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))

	y := func(l int, b int) int {
		return l * b
	} (20, 30)
	fmt.Printf("%T \n", y)
	fmt.Println(y)

	z := func(s string) string {
		return strings.ToUpper(s)
    }
    fmt.Printf("%T \n", z)
    fmt.Println(z("Joe"))

	cube := func(i int) string {
        c := i * i * i
        return strconv.Itoa(c)
    }
    z1 := cube(8)
    fmt.Printf("%T %v", z1, z1)

}