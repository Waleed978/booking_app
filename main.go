package main

import "fmt"

type Person struct {
	id   string
	name string
	age  int
}

func (p Person) getPerson() string {
	return fmt.Sprintf("Person: %s, %s, %d", p.id, p.name, p.age)
}

// type Car struct {
// 	model string
// 	year  int
// 	wheel struct {
// 		width  int
// 		height int
// 	}
// }
type Car struct {
	model string
	year  int
	wheel Wheel
}
type Wheel struct {
	width  int
	height int
}

func main() {
	// car := Car{model: "Toyota", year: 2020, wheel: struct {
	// 	width  int
	// 	height int
	// }{width: 10, height: 20}}

	// fmt.Println(car.model)
	// fmt.Println(car.year)
	// fmt.Println(car.wheel.width)
	car := Car{model: "Toyota", year: 2020, wheel: Wheel{width: 10, height: 20}}
	fmt.Println(car.wheel.height)
	fmt.Println(car.model)
	fmt.Println(car.year)
	fmt.Println(car.wheel.width)

	person := Person{id: "1", name: "Waleed", age: 25}
	fmt.Println(person.getPerson())
	name := "Waleed"
	age := 25

	fmt.Printf("My name is %s and I am %T years old.\n", name, age)
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, 10.12)
	fmt.Println(msg)
	if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}
	sum, sub := addAndSub(10, 20)
	fmt.Printf("The sum of %d and %d is %d\n", 10, 20, sum)
	fmt.Printf("The difference of %d and %d is %d\n", 10, 20, sub)

}
func addAndSub(a int, b int) (int, int) {
	return a + b, a - b
}
