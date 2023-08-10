package main
import (
	"fmt"
)

func addHundred(x int) int {
    return x + 100
}
func partialSum(add100 func(x int) int, x ...int) int {
    sum := 0
    for _, value := range x {
        sum += value
    }
    return add100(sum)

}

func calcArea(r float64) float64 {
    return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
    return 3.14 * r * 2
}

func calcDiameter(r float64) float64 {
    return r * 2
}

func getFunction(query int) func(r float64) float64 {

    queryToFunc := map[int]func(r float64) float64 {
        1: calcArea,
        2: calcPerimeter,
        3: calcDiameter,
    }
    return queryToFunc[query]
}

func printResult(radius float64, calcFunction func(r float64) float64) {
    result := calcFunction(radius)
    fmt.Println("Result: ", result)
}

func main() {

	partial := partialSum(addHundred, 1, 2, 3)
    fmt.Println(partial)

    var query int 
	var radius float64

    fmt.Printf("Enter the radius of the circle: ")
    fmt.Scanf("%f\n", &radius)

    fmt.Printf("Enter \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
    fmt.Scanf("%d\n", &query)

    printResult(radius, getFunction(query))
}