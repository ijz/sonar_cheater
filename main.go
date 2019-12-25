package main

import (
	"fmt"
	"sonar_cheater/gameplay"
	"sonar_cheater/terrains"
)

func makeMove(gb *gameplay.GameBoard, direction gameplay.MoveDirection) {
	var action gameplay.Action = &gameplay.MoveAction{Direction:direction}
	if err := gb.AcceptAction(&action); nil != err {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("ARCHIPELAGO")
	archipelago := terrains.MakeArchipelago()
	terrains.PrintTerrain(archipelago)

	gb := gameplay.NewGameBoard(archipelago)

	makeMove(gb, gameplay.MoveDirectionRight)
	//makeMove(gb, gameplay.MoveDirectionRight)
	makeMove(gb, gameplay.MoveDirectionDown)
	//makeMove(gb, gameplay.MoveDirectionLeft)
	//
	//makeMove(gb, gameplay.MoveDirectionDown)
	//makeMove(gb, gameplay.MoveDirectionDown)
	makeMove(gb, gameplay.MoveDirectionRight)
	//makeMove(gb, gameplay.MoveDirectionRight)
	//
	//makeMove(gb, gameplay.MoveDirectionRight)
	//makeMove(gb, gameplay.MoveDirectionRight)

	gb.PrintPath()

	//gb.RecalculateStartPoints(-1, -1)

	//var sonarAction gameplay.Action = &gameplay.SonarAction{Row:1}
	//gb.AcceptAction(&sonarAction)

	var silenceAction gameplay.Action = &gameplay.SilenceAction{}
	gb.AcceptAction(&silenceAction)

	makeMove(gb, gameplay.MoveDirectionRight)

	gb.PrintPath()

	//startPoint, err := gb.GetStartPoint()
	//if nil != err {
	//	fmt.Printf("Cannot get start point: %s", err)
	//	return
	//}
	startPoint := terrains.CombineUint8(0, 2)
	possibleLocations, err := gb.FindPossibleLocations(startPoint)
	for _, i := range possibleLocations {
		fmt.Printf("Possible location: %s\n", terrains.StringUint16(i))
	}
	fmt.Println(err)
}
