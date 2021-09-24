package main

import (
	"fmt"
	"hello/world"
)

func main() {
	sayHi := world.Hi("vigo")
	sayBye := world.Bye("vigo")

	fmt.Println(sayHi)  // Hi vigo!
	fmt.Println(sayBye) // Bye vigo!
}
