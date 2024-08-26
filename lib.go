// Package ex can perform some simple math operations
// Part of learning go ch10 exercises
package ex

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two integers
// This is a simple function that adds two integers
// and returns the result
// Example:
//
//	Add(1, 2) // 3
//
// [Addition doc]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
