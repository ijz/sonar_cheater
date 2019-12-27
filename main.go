package main

import (
	"sonar_cheater/gameplay"
)



func main() {
	ki := gameplay.NewKeyboardInput()
	ki.GetTerrain()

	ki.Loop()
}
