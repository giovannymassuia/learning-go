package math

// visible only inside package
func sum[T int | float64](a, b T) T {
	return a + b
}

// visible outside of package
func Sum[T int | float64](a, b T) T {
	return a + b
}

// visible outside of package: starts with capital letter
// visible only inside package: starts with lowercase letter
