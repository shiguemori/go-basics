package main

import "fmt"

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator
}

// a new reptile will be created of the same species that laid the egg, can't hatch twice
func (egg *ReptileEgg) Hatch() Reptile {
	if egg.CreateReptile != nil {
		childDragon := egg.CreateReptile()
		egg.CreateReptile = nil
		return childDragon
	}
	return nil
}

// The FireDragon type implements the Reptile interface
type FireDragon struct {
}

func (dragon *FireDragon) Lay() ReptileEgg {
	return ReptileEgg{
		CreateReptile: func() Reptile {
			return &FireDragon{}
		},
	}
}

func main() {
	var fireDragon FireDragon
	var egg ReptileEgg = fireDragon.Lay()
	var childDragon Reptile = egg.Hatch()
	var childDragon2 Reptile = egg.Hatch()
	fmt.Println(childDragon,childDragon2)
}