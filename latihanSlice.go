package main

import (
	"fmt"
)

func main() {
	//append n copy
	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "purple")
	// fmt.Println(colors)

	copied_colors := make([]string, 20)

	copy(copied_colors, colors)
	fmt.Println(copied_colors)
}
