package main

// type assertion

func main() {

	var myVar interface{} = "foo"

	println(myVar)
	println(myVar.(string))

	res, ok := myVar.(string)
	println(res, ok)

	res2, ok2 := isString(myVar)
	println(res2, ok2)

	res3, ok3 := myVar.(int)
	println(res3, ok3)

}

func isString(param interface{}) (int, bool) {
	res, ok := param.(int)
	return res, ok
}
