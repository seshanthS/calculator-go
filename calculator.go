package main

import ("fmt")

func home(){
	fmt.Printf(" 1.Addition \n 2.Subtraction \n 3.Average \n")
	fmt.Printf("Enter the number for the corresponding Operation : ")
	test := make(map[int]string)
	test[1]="Addition"
	test[2]="subtraction"
	test[3]="Average"
	var choice int
	fmt.Scanf("%d",&choice)
	fmt.Printf("\n your choice is %s \n",test[choice])
	
	//act on the choice
	
}	
	
func main(){
	home()
}
