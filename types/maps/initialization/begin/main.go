// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// initialize the map with make
	authors := map[string]author{
		"tm": author{name: "Toni Morrison"},
		"ma": author{name: "Marcus Aurelius"},
	}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	fmt.Println(authors["tm"])
}
