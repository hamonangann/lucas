// Package lucas is a library for generating Lucas sequence and counting Nth number of Lucas sequence.
// There are many sequences supported (see example.go)
//
// - Fibonacci numbers
// - Lucas numbers
// - Pell numbers
// - Counting numbers
// - Jacobsthal numbers
// - Generalized Lucas sequences by setting two (P and Q) parameters
//
// This library supports big numbers and return big.Int pointer(s)
//
// Example:
//
//	fib9 := lucas.GenerateFibonacci(10)
//	fmt.Println(fib9[9].Int64()) // 34
//
//	fib10 := lucas.NthFibonacci(10)
//	fmt.Println(fib10.Int64()) // 55
//
// [Lucas sequence]: https://en.wikipedia.org/wiki/Lucas_sequence
package lucas
