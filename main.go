package main

import (
	"fmt"
	"sonar_cheater/terrains"
)

func main() {
	fmt.Println("ARCHIPELAGO")
	archipelago := terrains.MakeArchipelago()
	terrains.PrintTerrain(archipelago)

	fmt.Println("VOLCANO")
	volcano := terrains.MakeVolcano()
	terrains.PrintTerrain(volcano)

	fmt.Println("NORTH POLE")
	northPole := terrains.MakeNorthPole()
	terrains.PrintTerrain(northPole)

	fmt.Println("OPEN WATERS")
	openWaters := terrains.MakeOpenWaters()
	terrains.PrintTerrain(openWaters)
}
