package gameplay

import "errors"

type Action interface {
	EnergyCost() int8
	// TODO: change current state type
	Perform(currentState int) error
}

type MoveDirection uint8
const (
	MoveDirectionUp MoveDirection = 0
	MoveDirectionDown MoveDirection = 1
	MoveDirectionLeft MoveDirection = 2
	MoveDirectionRight MoveDirection = 3
)

type MoveAction struct {
	direction MoveDirection
}

func (m MoveAction) EnergyCost() int8 {
	return -1
}

func (m MoveAction) isAllowed(currentState int) bool {
	// TODO: to implement
	if m.direction == MoveDirectionUp {
		return true
	}
	return false
}

func (m MoveAction) Perform(currentState int) error {
	if !m.isAllowed(currentState) {
		return errors.New("not allowed")
	}
	return nil
}
