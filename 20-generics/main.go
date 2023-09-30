package main

// generics

type MyNumber int

func sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// Number constraints
type Number interface {
	~int | float64
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	println(sum(m))

	m2 := map[string]float64{"foo": 1.1, "bar": 2.2, "baz": 3.3}
	println(sum(m2))

	m3 := map[string]MyNumber{"foo": 1, "bar": 2, "baz": 3}
	println(sum(m3))

	println(Compare(1, 1))
	println(Compare(1, 2))
	println(Compare(1, 1.0))
	//println(Compare(1, "1")) // doesn't work
}
