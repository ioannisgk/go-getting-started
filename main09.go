package main
import (
	"fmt"
)

type Movie struct {
    name   string
    rating float32
}

func getMovie(s string, r float32) (m Movie) {
    m = Movie{
        name:   s,
        rating: r,
    }
    return
}

func increaseRating(m *Movie) {
    m.rating += 1.0
}

func main() {

    mov1 := getMovie("xyz", 2.0)
    mov2 := getMovie("abc", 3.3)
    increaseRating(&mov1)
    fmt.Printf("%+v", mov1)
    fmt.Printf("%+v", mov2)
    
    movies := make([]Movie, 1)
    movies = append(movies, mov1)
    movies = append(movies, mov2)
    for _, value := range movies {
        fmt.Println(value)
    }
    
    if mov1.rating == mov2.rating || mov1 != mov2 {
        fmt.Println("Movies are different")
    } else if mov1.rating == mov2.rating {
        fmt.Println("Movies have the same rating")
    }

}