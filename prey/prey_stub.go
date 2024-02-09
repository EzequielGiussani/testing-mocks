package prey

import "testdoubles/positioner"

// NewTuna creates a new Tuna
func NewPreyStub() Prey {
	return &PreyStub{}
}

// Tuna is an implementation of the Prey interface
type PreyStub struct {
	FuncGetSpeed    func() (speed float64)
	FuncGetPosition func() (position *positioner.Position)
}

func (s *PreyStub) GetSpeed() (speed float64) {
	speed = s.FuncGetSpeed()
	return
}

func (s *PreyStub) GetPosition() (position *positioner.Position) {
	position = s.FuncGetPosition()
	return
}
