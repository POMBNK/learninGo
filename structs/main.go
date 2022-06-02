package main

import "fmt"

type Gun struct {
	On    bool
	Ammo  int
	Power int
}

func (g *Gun) Shoot() bool {
	if g.Ammo >= 1 && g.On == true {
		g.Ammo--
		return true
	} else {
		return false
	}
}

func (g *Gun) RideBike() bool {
	if g.Power >= 1 && g.On == true {
		g.Power--
		return true
	} else {
		return false
	}
}

func main() {
	testStruct := &Gun{false, 10, 1}
	fmt.Println(testStruct.Shoot(), testStruct.RideBike())
}
