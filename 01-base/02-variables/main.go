package main

// constant, immutable, final
const a = "Hello"

// variable in global scope
var b bool // default if false
var (
	c int
	d string
	e float64
)

func main() {
	// local scope
	var a string

	// shorthand
	// declare, attribute
	// `:=` only first time, after only `=`
	f := "string"

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
