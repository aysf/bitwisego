package main

/*
This example demonstrates bitwise operator usage to find out whether
there is a certain medal in Ash's medal collection
*/

import "fmt"

const (
	fire  = 1 << iota // 0001
	water             // 0010
	wind              // 0100
	earth             // 1000
)

func main() {

	// try to input other value combination
	AshMedal := fire | water | wind

	// uncommented this value to add earth medal to the collection
	// AshMedal ^= earth

	// check if there is any earth medal in ash's medal collection
	if AshMedal&earth != 0 {
		fmt.Println("Ash has earth medal")
	} else {
		fmt.Println("Ash has no earth medal")
	}

}
