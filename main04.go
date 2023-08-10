package main
import (
	"fmt"
)

func main() {

    map1 := make(map[string]int)
	fmt.Println(map1)

	map2 := map[string]string{
	    "en": "English",
	    "fr": "French" }
	fmt.Println(map2)
	fmt.Println(len(map2))

	fmt.Println(map2["en"])
	fmt.Println(map2["fr"])

	value1, found1 := map2["en"]
	fmt.Println(found1, value1)

	value2, found2 := map2["it"]
	fmt.Println(found2, value2)

	map2["it"] = "Italian"
	fmt.Println(map2)

	delete(map2, "en")
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Println(key, "=>", value)
	}

	map2 = make(map[string]string)
	fmt.Println(map2)

	ascii_codes := map[string]int{}
    ascii_codes["A"] = 65
    _, found := ascii_codes["B"]
    if !(found) {
            fmt.Println("key B was not found")
    }

}