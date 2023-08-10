package main
import (
	"fmt"
	"reflect"
	"strconv"
)
/*  
    ======== OPERATORS ========

    Comparison: == != < > <= >=
	Arithmetic: + - * / % ++ --
	Assignment: = += -= *= /= %=
	Bitwise:    & | << >> ^
	Logical:    && || !
*/

var myGlobalVar string = "20"

func main() {

    user := "Adam"
	const place = "Greece"

    var (
		precision float64
	    grade int = 82
	)

	i, err := strconv.Atoi(myGlobalVar)

	fmt.Printf("Here is %v \n", place)
	fmt.Printf("My grade: %d \n", grade)
	fmt.Printf("My grade: %q \n", strconv.Itoa(grade))
	fmt.Printf("My version: %v, %v", i, err)

	fmt.Printf("\nHello %v, your grade is %.2f", user, float64(grade))
	fmt.Printf("\nThe type of %v is %T", precision, precision)
	fmt.Printf("\nType %v ", reflect.TypeOf(grade))

	fmt.Print("\nEnter your name: ")
	fmt.Scanf("%s", &user)
	fmt.Print("Hey there, ", user)

	var a, b float64 = 24.4, 3.0
    fmt.Println(a / b)
    fmt.Println(int(a) % int(b))

	var c bool = false
    result := 10 > 50
    fmt.Println(!(c && result))
}