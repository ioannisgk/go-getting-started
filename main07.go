package main
import (
	"fmt"
)

func modify1(s string) {
	s = "world"
}

func modify2(s *string) {
	*s = "world"
}

func modify3(s []int) {
	s[0] = 100
}

func modify4(m map[string]int) {
	m["K"] = 75
}

func modify(numbers ...int) {
    for i := range numbers {
        numbers[i] -= 5
    }
}

func main() {

	i := 10
	j := "hello"
	k := "world"

	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	var ptr_i *int = &i
	var ptr_j = &j
	ptr_k := &k

	fmt.Println(ptr_i)
	fmt.Println(ptr_j)
	fmt.Println(ptr_k)

	fmt.Printf("%T %v \n", j, j)
	*ptr_j = "HELLO"
	fmt.Printf("%T %v \n", j, j)

	a := "hello"
	fmt.Println(a)
	modify1(a)
	fmt.Println(a)

	modify2(&a)
	fmt.Println(a)

	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modify3(slice)
	fmt.Println(slice)

	asciiCodes := make(map[string]int)
	asciiCodes["A"] = 65
	asciiCodes["F"] = 70
	fmt.Println(asciiCodes)
	modify4(asciiCodes)
	fmt.Println(asciiCodes)

	arr := []int{10, 20, 30}
    fmt.Println(arr)
    modify(arr...)
    fmt.Println(arr)
}