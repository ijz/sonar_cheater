package gameplay

type MoveDirection uint8
const (
	MoveDirectionUp MoveDirection = 0
	MoveDirectionDown MoveDirection = 1
	MoveDirectionLeft MoveDirection = 2
	MoveDirectionRight MoveDirection = 3
	MoveDirectionSilent MoveDirection = 4
)

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
	return nil
}

type SonarAction struct {
	isRow bool
	location uint8
}

func (s *SonarAction) EnergyCost() int8 {
	return 2
}

func (s *SonarAction) Perform(gb *GameBoard) error {
	return nil
}

type SilenceAction struct {}

func (s *SilenceAction) EnergyCost() int8 {
	return 3
}

func (s *SilenceAction) Perform(gb *GameBoard) error {
	return nil
}

type TorpedoAction struct {
	location uint16
}

func (t *TorpedoAction) EnergyCost() int8 {
	return 4
}

func (t *TorpedoAction) Perform(gb *GameBoard) error {
	return nil
}

