package gameplay

import (
	"errors"
	"fmt"
	"sonar_cheater/terrains"
)

type Node struct {
	Action MoveDirection
	Children []*Node
}

func (n *Node) isCertain() bool {
	return 1 == len(n.Children)
}

type GameBoard struct {
	terrain []uint8
	possibleStartingPoints map[uint16]bool
	possibleLocations []uint16
	path *Node
	energy int8
	currentLeaves []*Node
	life uint8
}

func NewGameBoard(terrain []uint8) *GameBoard {
	var gb GameBoard
	gb.terrain = terrain
	gb.life = 2
	gb.possibleStartingPoints = make(map[uint16]bool)

	// all points are possible start points without moves
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			startPoint := terrains.CombineUint8(uint8(i), uint8(j))
			gb.possibleStartingPoints[startPoint] = true
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
		return 0, 0, errors.New("failed_move")
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

func (gb *GameBoard) FindPossibleLocations(startPoint uint16) ([]uint16, error) {
	if nil == gb.path {
		return nil, errors.New("no_path")
	}
	var validPositions []uint16
	err := gb.findPossibleLocationsHelper(startPoint, gb.path, &validPositions)
	return validPositions, err
}

func (gb *GameBoard) findPossibleLocationsHelper(
	startPoint uint16, node *Node, validPositions *[]uint16) error {
	if nil == node {
		*validPositions = append(*validPositions, startPoint)
		return nil
	}
	r, c := terrains.SplitUint16(startPoint)
	newr, newc, err := gb.move(r, c, node)
	if nil != err {
		return err
	}
	if nil == node.Children || 0 == len(node.Children) {
		*validPositions = append(*validPositions, terrains.CombineUint8(newr, newc))
		return nil
	}

	for _, c := range node.Children {
		// TODO: prune tree base off of err
		if err = gb.findPossibleLocationsHelper(terrains.CombineUint8(newr, newc), c, validPositions); nil != err {
			return err
		}
	}
	return nil
}

func (gb *GameBoard) RecalculateStartPoints(row int8, col int8) {
	// TODO debug
	for start, _ := range gb.possibleStartingPoints {
		locations, err := gb.FindPossibleLocations(start)
		startString := terrains.StringUint16(start)
		if nil != err {
			fmt.Printf("removing %s\n", startString)
			delete(gb.possibleStartingPoints, start)
		} else {
			fmt.Printf("possible start point %s\n", startString)
			for l := range locations {
				fmt.Printf("\tpossible current position %s\n", terrains.StringUint16(uint16(l)))
			}
		}

	}
}

func (gb *GameBoard) PrintPath() {
	var stack []*Node
	stack = append(stack, gb.path)

	for nil != stack && 0 != len(stack) {
		var row []*Node
		for _, node := range stack {
			fmt.Printf("%s ", DirectionDict[node.Action])
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

func (gb *GameBoard) ListPossibleLocations() []uint16 {
	return nil
}

func (gb *GameBoard) TorpedoHit(location uint16, didHit bool) (bool, error) {
	if didHit {
		// TODO how do we indicate on path
		gb.possibleStartingPoints = make(map[uint16]bool)
		gb.possibleStartingPoints[location] = true
		gb.life -= 1
	} else {
		// TODO prune tree
	}
	return 0 == gb.life, nil
}
