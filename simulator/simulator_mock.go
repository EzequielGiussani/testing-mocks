package simulator

import (
	"github.com/stretchr/testify/mock"
)

func NewCatchSimulatorMock() CatchSimulator {
	return &CatchSimulatorMock{}
}

type CatchSimulatorMock struct {
	mock.Mock
	FuncCanCatch func(hunter, prey *Subject) (ok bool)
}

func (c *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (ok bool) {
	// args := c.Called(hunter, prey)
	// ok = args.Bool(0) && args.Bool(1)
	// return
	ok = c.FuncCanCatch(hunter, prey)

	return
}
