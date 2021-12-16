package math

import "constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Add[T Number](a, b T) T {
	return a + b
}

func Sub[T Number](a, b T) T {
	return a - b
}

func Mul[T Number](a, b T) T {
	return a * b
}

func Div[T Number](a, b T) T {
	return a / b
}
