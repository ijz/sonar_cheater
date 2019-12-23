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
	possibleStartingPoints []uint16
	path *Node
	energy int8
	currentLeaves []*Node
	damage uint8
}

func NewGameBoard(terrain []uint8) *GameBoard {
	var gb GameBoard
	gb.terrain = terrain
	gb.possibleStartingPoints = make([]uint16, 100, 100)

	// all points are possible start points without moves
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			startPoint := terrains.CombineUint16(uint8(i), uint8(j))
			index := i * 10 + j
			gb.possibleStartingPoints[index] = startPoint
		}
	}
	return &gb
}

func (gb *GameBoard) AcceptAction(a *Action) error {
	gb.energy -= (*a).EnergyCost()
	if gb.energy < 0 {
		return errors.New("not_enough_energy")
	}
	if err := (*a).Perform(gb); nil != err {
		return err
	}
	return nil
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

func (gb *GameBoard) recalculateStartPoints(row int8, col int8) {

}

func (gb *GameBoard) PrintPath() {
	var stack []*Node
	stack = append(stack, gb.path)

	for nil != stack && 0 != len(stack) {
		var row []*Node
		for _, node := range stack {
			fmt.Printf("%d ", node.Action)
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
