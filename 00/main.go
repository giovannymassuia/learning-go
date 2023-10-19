package main

// Composition is a way to reuse code without using inheritance.
// Go does not support inheritance, but it does support composition.

type Animal struct {
}

func (Animal) Speak() {
	println("animal speaks")
}

type Dog struct {
	Animal
}

func (Dog) Bark() {
	println("dog speaks")
}

func main() {

	animal := Animal{}
	animal.Speak()

	dog := Dog{}
	dog.Speak()
	dog.Bark()

}
