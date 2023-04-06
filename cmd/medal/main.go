package main

/*
This example of using bitwise operator to find out whether there is
a certain medal in Ash's medal collection

*/

import "fmt"

const (
	fire  = 1 << iota // 0001
	water             // 0010
	wind              // 0100
	earth             // 1000
)

func main() {

	// commented/uncommented these values and
	// also try to input other values
	AshMedal := fire | water | wind
	// AshMedal := earth | fire

	// check earth medal in ash 'smedal collection
	if AshMedal&earth != 0 {
		fmt.Println("Ash has earth medal")
	} else {
		fmt.Println("Ash has no earth medal")
	}

}
