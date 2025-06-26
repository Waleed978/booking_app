package main

import "fmt"

func main() {
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
