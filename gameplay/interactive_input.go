package gameplay

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sonar_cheater/terrains"
	"strconv"
	"strings"
)

var CommandDict = map[string]func(*KeyboardInput, string) {
	"move": move,
	"sonar": sonar,
	"silence": silence,
	"surface": surface,
	"take_hit": takeHit,
	"hit": hit,
	"miss": miss,
	"quit": quit,
	"exit": quit,
}

var DirectionNameDict = map[string]MoveDirection {
	"up": MoveDirectionUp,
	"down": MoveDirectionDown,
	"left": MoveDirectionLeft,
	"right": MoveDirectionRight,
}

func move(ki *KeyboardInput, direction string) {
	if _, ok := DirectionNameDict[direction]; !ok {
		log.Printf("Invalid direction %s", direction)
		return
	}
	var action Action = &MoveAction{Direction:DirectionNameDict[direction]}
	if err := ki.gb.AcceptAction(&action); nil != err {
		log.Printf("error while moving: %s", err)
	}
}

func sonar(ki *KeyboardInput, location string) {
	row, col := terrains.Int8sFromString(location)
	var sonarAction Action = &SonarAction{Row: row, Col: col}
	if err := ki.gb.AcceptAction(&sonarAction); nil != err {
		log.Printf("error while sonaring: %s", err)
	}
}

func silence(ki *KeyboardInput, _ string) {
	var silenceAction Action = &SilenceAction{}
	if err := ki.gb.AcceptAction(&silenceAction); nil != err {
		log.Printf("error while silencing: %s", err)
	}
}

func surface(ki *KeyboardInput, location string) {
	row, col := terrains.Int8sFromString(location)
	l := terrains.CombineUint8(uint8(row), uint8(col))
	var surfaceAction Action = &SurfaceAction{location:l}
	if err := ki.gb.AcceptAction(&surfaceAction); nil != err {
		log.Printf("error while surfacing: %s", err)
	}
}

func takeHit(ki *KeyboardInput, location string) {
	row, col := terrains.Int8sFromString(location)
	l := terrains.CombineUint8(uint8(row), uint8(col))
	if err := ki.gb.TakeTorpedoHit(l); nil != err {
		log.Printf("error while hitting: %s", err)
	}
	ki.gb.PrintState()
}

func hit(ki *KeyboardInput, location string) {
	row, col := terrains.Int8sFromString(location)
	l := terrains.CombineUint8(uint8(row), uint8(col))
	if err := ki.gb.TorpedoHit(l); nil != err {
		log.Printf("error while hitting: %s", err)
	}
	ki.gb.PrintState()
}

func miss(ki *KeyboardInput, location string) {
	row, col := terrains.Int8sFromString(location)
	l := terrains.CombineUint8(uint8(row), uint8(col))
	ki.gb.TorpedoMiss(l)
	ki.gb.PrintState()
}

func quit(ki *KeyboardInput, _ string) {
	ki.shouldStop = true
}

func makeTerrain(terrain terrains.Terrain) []uint8 {
	t := terrains.TerrainDict[terrain]()
	fmt.Printf("%s\n", terrains.TerrainNameDict[terrain])
	terrains.PrintTerrain(t)
	return t
}


type KeyboardInput struct {
	reader *bufio.Reader
	gb *GameBoard
	shouldStop bool
}

func NewKeyboardInput() *KeyboardInput {
	ki := new(KeyboardInput)
	ki.reader = bufio.NewReader(os.Stdin)
	ki.shouldStop = false
	return ki
}

func (ki *KeyboardInput) readLine() (string, error) {
	if s, err := ki.reader.ReadString('\n'); nil == err {
		return strings.TrimSpace(s), nil
	} else {
		return "", err
	}
}

func (ki *KeyboardInput) Prompt(msg string) (string, error) {
	fmt.Printf("%s...", msg)
	return ki.readLine()
}

func (ki *KeyboardInput) GetTerrain() {
	t, err := ki.Prompt("Enter terrain [0 - 3]")
	if nil != err {
		log.Fatalf("Cannot get terrain: %s", err)
	}
	tries := 0
	for tries < 3 {
		tries++
		i, err := strconv.Atoi(t)
		if nil != err {
			log.Print(err)
			continue
		}
		if i < int(terrains.ArchipelagoTerrain) || i > int(terrains.OpenWatersTerrain) {
			fmt.Print("number must be between 0 - 3")
			continue
		}
		terrain := terrains.Terrain(i)
		ki.gb = NewGameBoard(makeTerrain(terrain))
		break
	}
	if nil == ki.gb {
		log.Fatal("Cannot get terrain")
	}
}

func (ki *KeyboardInput) Loop() {

	for !ki.shouldStop {
		c, err := ki.Prompt("Enter command")
		if nil != err {
			log.Printf("cannot read command: %s", err)
			continue
		}
		cSlice := strings.Fields(c)
		if 0 == len(cSlice) {
			fmt.Printf("invalid command parameters '%s'", c)
			continue
		}

		commandFunc, ok := CommandDict[cSlice[0]]
		if !ok {
			fmt.Printf("invalid command '%s'", cSlice[0])
			continue
		}
		if 1 == len(cSlice) {
			// dummy entry for commands without a parameter
			cSlice = append(cSlice, "")
		}
		commandFunc(ki, cSlice[1])
	}
}

