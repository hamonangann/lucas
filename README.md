# lucas

Go package to generate and count Nth number of Lucas sequence (Fibonacci, Lucas, Pell, etc.s). Nth counting is optimized

This library supports big numbers and return big.Int pointer(s).

## Example

```go
fib9 := lucas.GenerateFibonacci(10)
fmt.Println(fib9[9].Int64()) // 34
```

```go
fib10 := lucas.NthFibonacci(10)
fmt.Println(fib10.Int64()) // 55
```

## Installation

```shell
go get github.com/hamonangann/lucas
```

## Contribution

Fork [this repository](https://github.com/hamonangann/lucas), then open a [pull request](https://github.com/hamonangann/lucas/pulls) to improve this project. I welcome any kind of contributions!

If you have any questions, please raise an [issue](https://github.com/hamonangann/lucas/issues).