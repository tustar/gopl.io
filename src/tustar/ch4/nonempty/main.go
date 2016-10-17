package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	fmt.Println("------------")
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data)
	fmt.Println("------------")
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("%v\n", s)
	fmt.Println(remove(s, 2))
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
