package lucas

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	t.Run("should generate correct Fibonacci numbers", func(t *testing.T) {
		actual := GenerateFibonacci(5)
		expected := []*big.Int{
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(1),
			big.NewInt(2),
			big.NewInt(3),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth Fibonacci", func(t *testing.T) {
		assert.Equal(t, big.NewInt(55), NthFibonacci(10))
	})

	t.Run("should generate correct Lucas numbers", func(t *testing.T) {
		actual := GenerateLucas(5)
		expected := []*big.Int{
			big.NewInt(2),
			big.NewInt(1),
			big.NewInt(3),
			big.NewInt(4),
			big.NewInt(7),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth Lucas", func(t *testing.T) {
		assert.Equal(t, big.NewInt(123), NthLucas(10))
	})

	t.Run("should generate correct Pell numbers", func(t *testing.T) {
		actual := GeneratePell(5)
		expected := []*big.Int{
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(2),
			big.NewInt(5),
			big.NewInt(12),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth Pell", func(t *testing.T) {
		assert.Equal(t, big.NewInt(2378), NthPell(10))
	})

	t.Run("should generate correct Pell-Lucas numbers", func(t *testing.T) {
		actual := GeneratePellLucas(5)
		expected := []*big.Int{
			big.NewInt(2),
			big.NewInt(2),
			big.NewInt(6),
			big.NewInt(14),
			big.NewInt(34),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth Pell-Lucas", func(t *testing.T) {
		assert.Equal(t, big.NewInt(6726), NthPellLucas(10))
	})

	t.Run("should generate correct counting numbers", func(t *testing.T) {
		actual := GenerateCounting(5)
		expected := []*big.Int{
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(2),
			big.NewInt(3),
			big.NewInt(4),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth counting", func(t *testing.T) {
		assert.Equal(t, big.NewInt(100000), NthCounting(100000))
	})

	t.Run("should generate correct Jacobsthal numbers", func(t *testing.T) {
		actual := GenerateJacobsthal(5)
		expected := []*big.Int{
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(1),
			big.NewInt(3),
			big.NewInt(5),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth Jacobsthal", func(t *testing.T) {
		assert.Equal(t, big.NewInt(341), NthJacobsthal(10))
	})

	t.Run("should generate correct Jacobsthal-Lucas numbers", func(t *testing.T) {
		actual := GenerateJacobsthalLucas(5)
		expected := []*big.Int{
			big.NewInt(2),
			big.NewInt(1),
			big.NewInt(5),
			big.NewInt(7),
			big.NewInt(17),
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct Nth Jacobsthal-Lucas", func(t *testing.T) {
		assert.Equal(t, big.NewInt(1025), NthJacobsthalLucas(10))
	})

}
