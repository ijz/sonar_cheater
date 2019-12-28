package gameplay

import (
	"log"
	"sonar_cheater/terrains"
)

type MoveDirection uint8
const (
	MoveDirectionUp MoveDirection = 0
	MoveDirectionDown MoveDirection = 1
	MoveDirectionLeft MoveDirection = 2
	MoveDirectionRight MoveDirection = 3
)

var DirectionDict = map[MoveDirection]string {
	MoveDirectionUp: "Up",
	MoveDirectionDown: "Down",
	MoveDirectionLeft: "Left",
	MoveDirectionRight: "Right",
}

func cleanBoard(gb *GameBoard, sonarRow int8, sonarCol int8, notHere int16) {
	gb.RecalculateStartPoints(sonarRow, sonarCol, notHere)
	if startPoint, err1 := gb.GetStartPoint(); nil == err1 {
		validPos, err2 := gb.FindPossibleLocations(startPoint, true, sonarRow, sonarCol, notHere)
		if nil != err2 {
			log.Printf("failed to prune path: %s", err2)
			return
		}
		var validPosString []string
		for _, vp := range validPos {
			validPosString = append(validPosString, terrains.StringUint16(vp))
		}
		log.Printf("start point: %s -> %v", terrains.StringUint16(startPoint), validPosString)
	} else {
		// more than one possible start points
		for sp, _ := range gb.GetPossibleStartPoints() {
			validPos, err := gb.FindPossibleLocations(sp, false, sonarRow, sonarCol, notHere)
			if nil != err {
				log.Printf("failed to follow path: %s", err)
				continue
			}
			var validPosString []string
			for _, vp := range validPos {
				validPosString = append(validPosString, terrains.StringUint16(vp))
			}
			log.Printf("? start point: %s -> %v", terrains.StringUint16(sp), validPosString)
		}
	}
}

type Action interface {
	EnergyCost() int8
	Perform(gb *GameBoard) error
}

type MoveAction struct {
	Direction MoveDirection
}

func (m *MoveAction) EnergyCost() int8 {
	return -1
}

func (m *MoveAction) Perform(gb *GameBoard) error {
	if err := gb.AddNodes(m.Direction); nil != err {
		return err
	}
	cleanBoard(gb, -1, -1, -1)
	return nil
}

type SonarAction struct {
	Row int8
	Col int8
}

func (s *SonarAction) EnergyCost() int8 {
	return 0  // we sonared our enemies, don't cost them anything
}

func (s *SonarAction) Perform(gb *GameBoard) error {
	cleanBoard(gb, s.Row, s.Col, -1)
	gb.PrintState()
	return nil
}

type SilenceAction struct {}

func (s *SilenceAction) EnergyCost() int8 {
	return 3
}

func (s *SilenceAction) Perform(gb *GameBoard) error {
	err := gb.AddNodes(MoveDirectionUp, MoveDirectionDown, MoveDirectionLeft, MoveDirectionRight)
	if nil != err {
		return err
	}
	gb.PrintState()
	return nil
}

type SurfaceAction struct {
	location uint16
}

func (s *SurfaceAction) EnergyCost() int8 {
	return 0
}

func (s *SurfaceAction) Perform(gb *GameBoard) error {
	gb.possibleStartingPoints = make(map[uint16]bool)
	gb.possibleStartingPoints[s.location] = true
	gb.path = nil
	gb.PrintState()
	return nil
}
