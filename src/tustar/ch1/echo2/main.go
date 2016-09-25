package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Print(index)
		fmt.Print("=>")
		fmt.Println(arg)
	}
}
