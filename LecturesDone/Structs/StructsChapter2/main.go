package main

import "fmt"

type str string

// Method like we did for struct
func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str

	name = "Max"
	name.log()
}
