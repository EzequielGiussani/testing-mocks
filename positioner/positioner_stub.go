package positioner

func NewPositionerStub() (positionerStub *PositionerStub) {
	positionerStub = &PositionerStub{}
	return
}

type PositionerStub struct {
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

func (s *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = s.FuncGetLinearDistance(from, to)
	return
}
