package gameplay

import (
	"errors"
	"fmt"
	"log"
	"sonar_cheater/terrains"
)

func findInSlice(needle uint16, haystack []uint16) (int, bool) {
	for i, item := range haystack {
		if needle == item {
			return i, true
		}
	}
	return 0, false
}

type Node struct {
	Action MoveDirection
	Children []*Node
}

type GameBoard struct {
	terrain []uint8
	possibleStartingPoints map[uint16]bool
	possibleLocations []uint16
	path *Node
	energy int8
	currentLeaves []*Node
	torpedoHits []uint16
}

func NewGameBoard(terrain []uint8) *GameBoard {
	var gb GameBoard
	gb.terrain = terrain
	gb.possibleStartingPoints = make(map[uint16]bool)

	// all points are possible start points without moves
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			startPoint := terrains.CombineUint8(uint8(i), uint8(j))
			if !terrains.IsIsland(gb.terrain, uint8(i), uint8(j)) {
				gb.possibleStartingPoints[startPoint] = true
			}
		}
	}

	return &gb
}

func (gb *GameBoard) AcceptAction(a *Action) error {
	if gb.energy < (*a).EnergyCost() {
		return errors.New("not_enough_energy")
	}
	gb.energy -= (*a).EnergyCost()
	if err := (*a).Perform(gb); nil != err {
		return err
	}
	return nil
}

func (gb *GameBoard) move(r uint8, c uint8, node *Node) (uint8, uint8, error) {
	var newr, newc uint8
	switch node.Action {
	case MoveDirectionUp:
		// already at the top row
		if 0 == r {
			return 0, 0, errors.New("failed_move")
		}
		newr = r - 1
		newc = c
		break
	case MoveDirectionDown:
		// already at the bottom row
		if 9 == r {
			return 0, 0, errors.New("failed_move")
		}
		newr = r + 1
		newc = c
		break
	case MoveDirectionLeft:
		// already at the top row
		if 0 == c {
			return 0, 0, errors.New("failed_move")
		}
		newr = r
		newc = c - 1
		break
	case MoveDirectionRight:
		// already at the top row
		if 9 == c {
			return 0, 0, errors.New("failed_move")
		}
		newr = r
		newc = c + 1
		break
	default:
		return 0, 0, errors.New("bad_direction")
	}
	if terrains.IsIsland(gb.terrain, newr, newc) {
		return 0, 0, errors.New("hit_island")
	}
	return newr, newc, nil
}

func (gb *GameBoard) AddNodes(directions ... MoveDirection) error {
	numNodes := len(directions)
	if nil == gb.path {
		if 1 != numNodes {
			return errors.New("first_node_must_be_singular")
		}
		gb.path = &Node{directions[0], nil}
		gb.currentLeaves = make([]*Node, 1)
		gb.currentLeaves[0] = gb.path
		return nil
	}

	var newLeaves []*Node

	for _, leaf := range gb.currentLeaves {
		leaf.Children = make([]*Node, numNodes, numNodes)
		for i, dir := range directions {
			leaf.Children[i] = &Node{dir, nil}
			newLeaves = append(newLeaves, leaf.Children[i])
		}
	}
	gb.currentLeaves = newLeaves
	return nil
}

func (gb *GameBoard) FindPossibleLocations(startPoint uint16, shouldTrim bool,
	sonarRow int8, sonarCol int8, notHere int16,) ([]uint16, error) {
	if nil == gb.path {
		return nil, errors.New("no_path")
	}
	var validPositions []uint16
	var seenLocs []uint16
	err := gb.findPossibleLocationsHelper(
		startPoint, gb.path, &validPositions, &seenLocs, 0, shouldTrim, sonarRow, sonarCol, notHere)

	return validPositions, err
}

func (gb *GameBoard) findPossibleLocationsHelper(
	startPoint uint16, node *Node, validPositions *[]uint16,
	seenLocs *[]uint16, step int, shouldTrim bool,
	sonarRow int8, sonarCol int8, notHere int16,
	) error {

	*seenLocs = (*seenLocs)[0:step]
	if _, found := findInSlice(startPoint, *seenLocs); found {
		return errors.New("own_path")
	}
	*seenLocs = append(*seenLocs, startPoint)
	step++

	if nil == node {
		r, c := terrains.SplitUint16(startPoint)
		if 0 <= sonarRow &&  r != uint8(sonarRow) {
			return errors.New("sonar_row_mismatch")
		}
		if 0 <= sonarCol &&  c != uint8(sonarCol) {
			return errors.New("sonar_col_mismatch")
		}
		if 0 <= notHere && startPoint == uint16(notHere) {
			return errors.New("torpedo_miss_at_point")
		}
		*validPositions = append(*validPositions, startPoint)
		return nil
	}
	r, c := terrains.SplitUint16(startPoint)
	startString := terrains.StringUint16(startPoint)
	newr, newc, err := gb.move(r, c, node)
	newStart := terrains.CombineUint8(newr, newc)
	newStartString := terrains.StringUint16(newStart)

	if nil != err {
		return err
	}

	if nil == node.Children || 0 == len(node.Children) {
		return gb.findPossibleLocationsHelper(
			newStart, nil, validPositions, seenLocs, step, shouldTrim, sonarRow, sonarCol, notHere)
	}
	newChildren := make([]*Node, 0)
	for _, child := range node.Children {
		err = gb.findPossibleLocationsHelper(
			newStart, child, validPositions, seenLocs, step, shouldTrim, sonarRow, sonarCol, notHere)
		if nil != err {
			log.Printf(
				"error while locating. start: %s, move: %s, newStart: %s, newMove: %s, err: %s",
				startString, DirectionDict[node.Action], newStartString, DirectionDict[child.Action], err)
		} else {
			newChildren = append(newChildren, child)
		}
	}
	if shouldTrim && len(newChildren) < len(node.Children) && len(node.Children) > 1 {
		// trim silence nodes
		node.Children = newChildren
	}
	if 0 == len(newChildren) {
		log.Printf("no valid move from %s", terrains.StringUint16(startPoint))
		return errors.New("all_subsequent_moves_failed")
	}
	return nil
}

func (gb *GameBoard) RecalculateStartPoints(row int8, col int8) {
	// if we know the start point for sure, no need to do anything
	if 1 == len(gb.possibleStartingPoints) {
		return
	}
	for start, _ := range gb.possibleStartingPoints {
		startString := terrains.StringUint16(start)

		locations, err := gb.FindPossibleLocations(start, false, row, col, -1)
		if nil != err {
			log.Printf("removing %s\n", startString)
			delete(gb.possibleStartingPoints, start)
		} else {
			log.Printf("possible start point %s\n", startString)
			for _, l := range locations {
				log.Printf("\tpossible current position %s\n", terrains.StringUint16(l))
			}
		}
	}
	log.Printf("All possible start points(%d):\n", len(gb.possibleStartingPoints))
	for start, _ := range gb.possibleStartingPoints {
		log.Println(terrains.StringUint16(start))
	}
}

func (gb *GameBoard) GetStartPoint() (uint16, error) {
	if 1 != len(gb.possibleStartingPoints) {
		return 0, errors.New("uncertain")
	}
	for l, _ := range gb.possibleStartingPoints {
		return l, nil
	}
	return 0, errors.New("impossible")
}

func (gb *GameBoard) PrintPath() {
	var stack []*Node
	stack = append(stack, gb.path)

	for nil != stack && 0 != len(stack) {
		var row []*Node
		for _, node := range stack {
			fmt.Printf("%s | ", DirectionDict[node.Action])
			if nil == node.Children {
				continue
			}
			for _, c := range node.Children {
				row = append(row, c)
			}
		}
		fmt.Println()
		stack = row
	}
}

func (gb *GameBoard) TorpedoHit(location uint16, didHit bool) error {
	if didHit {
		gb.possibleStartingPoints = make(map[uint16]bool)
		gb.possibleStartingPoints[location] = true
		gb.torpedoHits = append(gb.torpedoHits, location)
	} else {
		// TODO prune tree
	}
	if 2 == len(gb.torpedoHits) {
		fmt.Println("We are victorious!")
		// TODO print full path and hit
	}
	return nil
}
