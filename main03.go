package main
import (
	"fmt"
)

func main() {

	arrA := [5]bool{true, true, true, true}
    for i := 0; i < len(arrA); i++ {
        if arrA[i] {
            fmt.Println(arrA[i])
        }
    }
	fmt.Println(arrA[4])

    arrB := [5]int{90, 80, 70, 80, 97}
    for i := 0; i < len(arrB); i++ {
        fmt.Println(arrB[i])
    }

	arrC := [5]string{"a", "b", "c"}
    for index, element := range arrC {
        fmt.Println(index, "->", element)
    }

    sliceA := []int{10, 20, 30}
    for _, value := range sliceA {
        fmt.Println(value)
    }

    arrD := [5]string{"a", "b", "c", "d", "e"}
    sliceB := arrD[0:4]
    fmt.Println(arrD)
    fmt.Println(sliceB)

    sliceB[1] = "x"
    fmt.Println(arrD)
    fmt.Println(sliceB)

    arrE := [4]int{10, 20, 30, 40}
    sliceC := arrE[1:3]
    fmt.Println(arrE)
    fmt.Println(sliceC)

    sliceC = append(sliceC, 900, -90, 50)
    fmt.Println(arrE)
    fmt.Println(sliceC)

    arrF := [5]int{10, 20, 90, 70, 60}
    sliceD := arrF[0:3]
    fmt.Println(cap(sliceD))

    sliceE := make([]int, 5, 20)
    new_slice := append(sliceD, sliceE...)
    fmt.Println(cap(new_slice))

    arrG := [5]int{10, 20, 30, 40, 50}
    i := 2
    slice_1 := arrG[:i]
    slice_2 := arrG[i+1:]
    sliceF := append(slice_1, slice_2...)
    fmt.Println(sliceF)

    arrH := []int{10, 20, 90, 70, 60}
    sliceG := make([]int, 10)
    num := copy(sliceG, arrH)
    fmt.Println(sliceG)
    fmt.Println(num)

}