package simulator_test

import (
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimulator(t *testing.T) {

	t.Run("hunter catches the prey", func(t *testing.T) {
		// Arrange
		sim := simulator.NewCatchSimulatorMock()
		sim.CanCatchFunc = func(hunter, prey *simulator.Subject) (ok bool) {
			return true
		}

		// Act
		ok := sim.CanCatch(&simulator.Subject{}, &simulator.Subject{})

		// Assert
		require.True(t, ok)
		require.Equal(t, 1, sim.Calls.CanCatch)
	})

}
