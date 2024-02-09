package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPositioner(t *testing.T) {

	t.Run("should return a positioner with the correct values", func(t *testing.T) {
		// Arrange
		ps := positioner.NewPositionerStub()
		expectedDistance := float64(10)

		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = expectedDistance
			return
		}

		// Act
		distance := ps.GetLinearDistance(nil, nil)

		// Assert
		require.Equal(t, expectedDistance, distance)
	})

	t.Run("coordinates should all be negative", func(t *testing.T) {
		// Arrange
		ps := positioner.NewPositionerStub()
		expectedDistance := float64(1)

		from := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}

		to := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}

		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			if from.X < 0 && from.Y < 0 && from.Z < 0 && to.X < 0 && to.Y < 0 && to.Z < 0 {
				return 1
			}
			return 0
		}

		// Act
		distance := ps.GetLinearDistance(from, to)

		// Assert
		require.Equal(t, expectedDistance, distance)
	})

	t.Run("coordinates should all be positive", func(t *testing.T) {
		// Arrange
		ps := positioner.NewPositionerStub()
		expectedDistance := float64(1)

		from := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}

		to := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}

		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			if from.X > 0 && from.Y > 0 && from.Z > 0 && to.X > 0 && to.Y > 0 && to.Z > 0 {
				return 1
			}
			return 0
		}

		// Act
		distance := ps.GetLinearDistance(from, to)

		// Assert
		require.Equal(t, expectedDistance, distance)
	})

}
