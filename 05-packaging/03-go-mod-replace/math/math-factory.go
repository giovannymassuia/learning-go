package math

// NewMath can use the 'math' struct because it's in the same package, even if it's in a different file.
func NewMath(a, b int) *math {
	return &math{a, b}
}
