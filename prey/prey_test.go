package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrey(t *testing.T) {

	t.Run("Should return default values", func(t *testing.T) {
		// Arrange
		p := prey.NewPreyStub()
		p.GetSpeedFunc = func() (speed float64) {
			return 0
		}
		p.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{}
		}

		expectedSpeed := 0.0
		expectedPosition := &positioner.Position{}

		// Act
		speed := p.GetSpeed()
		position := p.GetPosition()

		// Assert
		require.Equal(t, expectedSpeed, speed)
		require.Equal(t, expectedPosition, position)
	})

	t.Run("Should return specified values", func(t *testing.T) {
		// Arrange
		p := prey.NewPreyStub()
		p.GetSpeedFunc = func() (speed float64) {
			return 10
		}
		p.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{
				X: 10,
				Y: 10,
				Z: 10,
			}
		}

		expectedSpeed := 10.0
		expectedPosition := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}

		// Act
		speed := p.GetSpeed()
		position := p.GetPosition()

		// Assert
		require.Equal(t, expectedSpeed, speed)
		require.Equal(t, expectedPosition, position)
	})

}
