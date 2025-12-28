package lucas

import (
	"math/big"
	"sync"
)

// GenerateSequenceU returns first n numbers of first kind of Lucas Sequence (U[0] to U[n-1])
// Choose this sequence to get starting U[0] = 0, U[1] = 1
// Choose P, Q so that U[n] = P*U[n-1] - Q*U[n-2]
// Complexity: O(n)
func GenerateSequenceU(p int64, q int64, n uint) []*big.Int {
	seq := make([]*big.Int, n)
	seq[0], seq[1] = big.NewInt(0), big.NewInt(1)

	// u(n) = p*u(n-1) - q*u(n-2)
	for i := uint(2); i < n; i++ {
		x, y, z := big.NewInt(0), big.NewInt(0), big.NewInt(0)
		x.Mul(big.NewInt(p), seq[i-1])
		y.Mul(big.NewInt(q), seq[i-2])
		seq[i] = z.Sub(x, y)
	}

	return seq
}

// GenerateSequenceV returns first n numbers of second kind of Lucas Sequence (V[0] to V[n-1])
// Choose this sequence to get starting V[0] = 2, V[1] = P
// Choose P, Q so that V[n] = P*V[n-1] - Q*V[n-2]
// Complexity: O(n)
func GenerateSequenceV(p int64, q int64, n uint) []*big.Int {
	seq := make([]*big.Int, n)
	seq[0], seq[1] = big.NewInt(2), big.NewInt(p)

	// v(n) = p*v(n-1) - q*v(n-2)
	for i := uint(2); i < n; i++ {
		x, y, z := big.NewInt(0), big.NewInt(0), big.NewInt(0)
		x.Mul(big.NewInt(p), seq[i-1])
		y.Mul(big.NewInt(q), seq[i-2])
		seq[i] = z.Sub(x, y)
	}

	return seq
}

var memo sync.Map

func countNthSequenceU(p int64, q int64, n uint) *big.Int {
	if v, ok := memo.Load(n); ok {
		return v.(*big.Int)
	}

	nHalfPlusOne := countNthSequenceU(p, q, (n/2)+1)
	nHalf := countNthSequenceU(p, q, n/2)

	x, y, z := big.NewInt(0), big.NewInt(0), big.NewInt(0)

	// Fast doubling formula for even n
	// u(2n) = 2*u(n)*u(n+1) - p*u(n)*u(n)
	if n%2 == 0 {
		x.Mul(big.NewInt(2), nHalf)
		x.Mul(x, nHalfPlusOne)
		y.Mul(big.NewInt(p), nHalf)
		y.Mul(y, nHalf)
		z.Sub(x, y)
	} else {
		// Fast doubling formula for odd n
		// u(2n+1) = u(n+1)*u(n+1) - q*u(n)*u(n)
		x.Mul(nHalfPlusOne, nHalfPlusOne)
		y.Mul(nHalf, nHalf)
		y.Mul(big.NewInt(q), y)
		z.Sub(x, y)
	}

	memo.Store(n, z)
	return z
}

// NthSequenceU returns n-th number of first kind of Lucas Sequence (zero-indexed)
// It is optimized with fast-doubling and thread-safe memoization to support larger n numbers
// Choose this sequence to get starting U[0] = 2, U[1] = P
// Choose P, Q so that U[n] = P*U[n-1] - Q*U[n-2]
// Complexity: O(log n)
func NthSequenceU(p int64, q int64, n uint) *big.Int {
	memo.Clear()
	memo.Store(uint(0), big.NewInt(0))
	memo.Store(uint(1), big.NewInt(1))
	memo.Store(uint(2), big.NewInt(p))
	return countNthSequenceU(p, q, n)
}

func exponent2N(x int64, n uint, xPowN *big.Int) *big.Int {
	res := big.NewInt(0).Mul(xPowN, xPowN)

	if n%2 == 1 {
		res.Mul(res, big.NewInt(x))
	}

	return res
}

func countNthSequenceV(p int64, q int64, n uint) (*big.Int, *big.Int) {
	if v, ok := memo.Load(n); ok {
		v0, _ := v.([2]*big.Int)
		return v0[0], v0[1]
	}

	qHalf, nHalf := countNthSequenceV(p, q, n/2)
	qNow := exponent2N(q, n/2, qHalf)

	x, y, z := big.NewInt(0), big.NewInt(0), big.NewInt(0)

	// Fast doubling formula for even n
	// v(2n) = v(n)*v(n) - 2*q(^n)
	if n%2 == 0 {
		x.Mul(nHalf, nHalf)
		y.Mul(big.NewInt(2), qNow)
		z.Sub(x, y)
	} else {
		// Fast doubling formula for odd n
		// v(2n+1) = v(n)*v(n+1) - p*q(^n)
		_, nHalfPlusOne := countNthSequenceV(p, q, (n/2)+1)
		x.Mul(nHalf, nHalfPlusOne)
		y.Mul(big.NewInt(p), qNow)
		z.Sub(x, y)
	}

	memo.Store(n, [2]*big.Int{qNow, z})
	return qNow, z
}

// NthSequenceV returns n-th number of second kind of Lucas Sequence (zero-indexed)
// It is optimized with fast-doubling, bin-exp, and thread-safe memoization to support larger n numbers
// Choose this sequence to get starting V[0] = 2, V[1] = P
// Choose P, Q so that V[n] = P*V[n-1] - Q*V[n-2]
// Complexity: O(log n)
func NthSequenceV(p int64, q int64, n uint) *big.Int {
	memo.Clear()
	memo.Store(uint(0), [2]*big.Int{big.NewInt(1), big.NewInt(2)})
	memo.Store(uint(1), [2]*big.Int{big.NewInt(1), big.NewInt(p)})
	_, res := countNthSequenceV(p, q, n)
	return res
}
