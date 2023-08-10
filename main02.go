package main
import (
	"fmt"
)

func main() {

    var a, b string = "foo", "bar"
    if a+b == "foo" {
        fmt.Println("foo")
    } else if a+b == "bar" {
         fmt.Println("bar")
    } else if a+b == "foobar" {
        fmt.Println("foobar")
    } else {
        fmt.Println("None matched")
    }

    switch a+b {
        case "foo":
            fmt.Println("foo")
        case "bar":
            fmt.Println("bar")
        case "foobar":
            fmt.Println("foobar")
        default:
            fmt.Println("None matched")
    }

    var i, j = 10, 50
    switch {
        case i+j == 60:
            fmt.Println("equal to 60")
        case i+j <= 60:
            fmt.Println("less than or equal to 60")
            fallthrough
        default:
            fmt.Println("greater than 60")
    }
	
    for i := 0; i <= 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i * i)
    }
}