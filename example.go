package lucas

import "math/big"

// GenerateFibonacci returns first n Fibonacci numbers (F[0] to F[n-1])
// e.g. [0 1 1 2 3 5 8 13 21 34 55]
func GenerateFibonacci(n uint) []*big.Int {
	return GenerateSequenceU(1, -1, n)
}

// NthFibonacci returns n-th Fibonacci number (zero-indexed)
// e.g. F[10] = 55
func NthFibonacci(n uint) *big.Int {
	return NthSequenceU(1, -1, n)
}

// GenerateLucas returns first n Lucas numbers (L[0] to L[n-1])
// e.g. [2 1 3 4 7 11 18 29 47 76 123]
func GenerateLucas(n uint) []*big.Int {
	return GenerateSequenceV(1, -1, n)
}

// NthLucas returns n-th Lucas number (zero-indexed)
// e.g. L[10] = 123
func NthLucas(n uint) *big.Int {
	return NthSequenceV(1, -1, n)
}

// GeneratePell returns first n Pell numbers (P[0] to P[n-1])
// e.g. [0 1 2 5 12 29 70 169 408 985 2378]
func GeneratePell(n uint) []*big.Int {
	return GenerateSequenceU(2, -1, n)
}

// NthPell returns n-th Pell number (zero-indexed)
// e.g. P[10] = 2378
func NthPell(n uint) *big.Int {
	return NthSequenceU(2, -1, n)
}

// GeneratePellLucas returns first n Pell-Lucas (companion Pell) numbers (Q[0] to Q[n-1])
// e.g. [2 2 6 14 34 82 198 478 1154 2786 6726]
func GeneratePellLucas(n uint) []*big.Int {
	return GenerateSequenceV(2, -1, n)
}

// NthPellLucas returns n-th Pell-Lucas (companion Pell) number (zero-indexed)
// e.g. Q[10] = 6726
func NthPellLucas(n uint) *big.Int {
	return NthSequenceV(2, -1, n)
}

// GenerateCounting returns first n counting numbers (0 to n-1)
// e.g. [0 1 2 3 4 5 6 7 8 9 10]
func GenerateCounting(n uint) []*big.Int {
	return GenerateSequenceU(2, 1, n)
}

// NthCounting returns n-th counting number (zero-indexed)
// e.g. N[10] = 10
func NthCounting(n uint) *big.Int {
	return NthSequenceU(2, 1, n)
}

// GenerateJacobsthal returns first n Jacobsthal numbers (J[0] to J[n-1])
// e.g. [0 1 1 3 5 11 21 43 85 171 341]
func GenerateJacobsthal(n uint) []*big.Int {
	return GenerateSequenceU(1, -2, n)
}

// NthJacobsthal returns n-th Jacobsthal number (zero-indexed)
// e.g. J[10] = 341
func NthJacobsthal(n uint) *big.Int {
	return NthSequenceU(1, -2, n)
}

// GenerateJacobsthalLucas returns first n Jacobsthal-Lucas (companion Jacobsthal) numbers (j[0] to j[n-1])
// e.g. [2 1 5 7 17 31 65 127 257 511 1025]
func GenerateJacobsthalLucas(n uint) []*big.Int {
	return GenerateSequenceV(1, -2, n)
}

// NthJacobsthalLucas returns n-th Jacobsthal-Lucas (companion Jacobsthal) number (zero-indexed)
// e.g. j[10] = 1025
func NthJacobsthalLucas(n uint) *big.Int {
	return NthSequenceV(1, -2, n)
}
