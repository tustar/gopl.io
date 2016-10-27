package main

import "tustar/ch6/geometry"

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Printf("%v\n", geometry.PathDistance(perim))
	fmt.Printf("%v\n", perim.Distance())
}
