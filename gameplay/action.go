package gameplay

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
	gb.RecalculateStartPoints(-1, -1)
	return nil
}

type SonarAction struct {
	Row int8
	Col byte
}

func (s *SonarAction) EnergyCost() int8 {
	return 2
}

func (s *SonarAction) Perform(gb *GameBoard) error {
	var row, col int8 = -1, -1
	if 0 < s.Row {
		row = s.Row - 1
	}
	if 0 < s.Col {
		col = int8(s.Col - 'A')
	}
	gb.RecalculateStartPoints(row, col)
	// TODO prune tree
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
	return nil
}

type TorpedoAction struct {
	location uint16
	didHit bool
}

func (t *TorpedoAction) EnergyCost() int8 {
	return 4
}

func (t *TorpedoAction) Perform(gb *GameBoard) error {
	return gb.TorpedoHit(t.location, t.didHit)
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
	return nil
}
