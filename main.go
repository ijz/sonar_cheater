package main

import (
	"fmt"
	"sonar_cheater/gameplay"
	"sonar_cheater/terrains"
)

func main() {
	fmt.Println("ARCHIPELAGO")
	archipelago := terrains.MakeArchipelago()
	terrains.PrintTerrain(archipelago)

	gb := gameplay.NewGameBoard(archipelago)
	var action gameplay.Action = &gameplay.MoveAction{Direction:gameplay.MoveDirectionRight}
	if err := gb.AcceptAction(&action); nil != err {
		fmt.Println(err)
	}

	action = &gameplay.MoveAction{Direction:gameplay.MoveDirectionDown}
	if err := gb.AcceptAction(&action); nil != err {
		fmt.Println(err)
	}

	action = &gameplay.MoveAction{Direction:gameplay.MoveDirectionLeft}
	if err := gb.AcceptAction(&action); nil != err {
		fmt.Println(err)
	}

	gb.PrintPath()

}
