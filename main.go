// demonstrate errors

package main

import (
	"fmt"
	"log"
)

func area(lenght float64, width float64) float64 {
	return lenght * width
}

func main() {
	lenght := 5.5
	width := 10

	result := area(lenght, float64(width))
	fmt.Println(result)
	log.Println(result)
}
