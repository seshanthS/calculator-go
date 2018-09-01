package main

import (
	"fmt"
)

func home() int{
	// UI
	fmt.Printf(" 1.Addition \n 2.Subtraction \n 3.Average \n 0. exit \n")
	fmt.Printf("Enter the number for the corresponding Operation : ")

	// Maps
	test := make(map[int]string)
	test[1] = "Addition"
	test[2] = "subtraction"
	test[3] = "Average"
	test[0] = "EXIT"

	var choice int
	fmt.Scanf("%d", &choice)
	fmt.Printf("\n your choice is %s \n", test[choice])

	//act on the choice
	switch choice {
	case 1:
		add()
	case 2:
		subtraction()
	case 3:
		ave()
	case 0:
	return 0
	default:
		home()
	}
	return 0
}

func main() {
	home()
	fmt.Println("exiting calculator...")

}

func subtraction() {
	var a, b int

	fmt.Printf("enter two numbers\n")
	fmt.Scanf("%d %d", &a, &b)
	sub := a - b
	fmt.Printf("the result is %d", sub)
}

func add() {

	var a, b int
	fmt.Println("check")
	fmt.Scanf("%d %d", &a, &b)

	c := a + b

	fmt.Printf("the addition is %d \n", c)

}
func ave() {

	var a, b, c int

	fmt.Printf("Enter three numbers\n")

	fmt.Scanf("%d %d %d", &a, &b, &c)

	ave := (a + b + c) / 3

	fmt.Printf("Average value is %d", ave)

}
