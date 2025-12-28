package lucas

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLucas(t *testing.T) {
	t.Run("should generate sequence U with length n", func(t *testing.T) {
		seq := GenerateSequenceU(1, -1, 10)

		assert.Len(t, seq, 10)
	})

	t.Run("should generate sequence U with correct first element", func(t *testing.T) {
		seq := GenerateSequenceU(1, -1, 10)

		assert.Equal(t, big.NewInt(0), seq[0])
	})

	t.Run("should generate sequence U with correct second element", func(t *testing.T) {
		seq := GenerateSequenceU(1, -1, 10)

		assert.Equal(t, big.NewInt(1), seq[1])
	})

	t.Run("should generate sequence U with correct recurrence relations", func(t *testing.T) {
		seq := GenerateSequenceU(1, -1, 10)
		actual := seq[len(seq)-1]
		expected := big.NewInt(0).Add(seq[len(seq)-2], seq[len(seq)-3])

		assert.Equal(t, expected, actual)
	})

	t.Run("should generate sequence V with length n", func(t *testing.T) {
		seq := GenerateSequenceV(1, -1, 10)

		assert.Len(t, seq, 10)
	})

	t.Run("should generate sequence V with correct first element", func(t *testing.T) {
		seq := GenerateSequenceV(1, -1, 10)

		assert.Equal(t, big.NewInt(2), seq[0])
	})

	t.Run("should generate sequence V with correct second element", func(t *testing.T) {
		seq := GenerateSequenceV(1, -1, 10)

		assert.Equal(t, big.NewInt(1), seq[1])
	})

	t.Run("should generate sequence V with correct recurrence relations", func(t *testing.T) {
		seq := GenerateSequenceV(1, -1, 10)
		actual := seq[len(seq)-1]
		expected := big.NewInt(0).Add(seq[len(seq)-2], seq[len(seq)-3])

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct first sequence U", func(t *testing.T) {
		first := NthSequenceU(1, -1, 0)

		assert.Equal(t, big.NewInt(0), first)
	})

	t.Run("should return correct second sequence U", func(t *testing.T) {
		second := NthSequenceU(1, -1, 1)

		assert.Equal(t, big.NewInt(1), second)
	})

	t.Run("should return correct third sequence U", func(t *testing.T) {
		third := NthSequenceU(1000, -1, 2)

		assert.Equal(t, big.NewInt(1000), third)
	})

	t.Run("should return correct Nth sequence U", func(t *testing.T) {
		nth := NthSequenceU(2, 1, 100000)

		assert.Equal(t, big.NewInt(100000), nth)
	})

	t.Run("should generate sequence V with correct recurrence relations", func(t *testing.T) {
		seq := GenerateSequenceV(1, -1, 10)
		actual := seq[len(seq)-1]
		expected := big.NewInt(0).Add(seq[len(seq)-2], seq[len(seq)-3])

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct first sequence V", func(t *testing.T) {
		first := NthSequenceV(1, -1, 0)

		assert.Equal(t, big.NewInt(2), first)
	})

	t.Run("should return correct second sequence V", func(t *testing.T) {
		second := NthSequenceV(1, -1, 1)

		assert.Equal(t, big.NewInt(1), second)
	})

	t.Run("should return correct Nth sequence V", func(t *testing.T) {
		nth := NthSequenceV(2, 1, 100000)

		assert.Equal(t, big.NewInt(2), nth)
	})
}
