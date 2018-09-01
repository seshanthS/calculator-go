package main



import "fmt"



func main() {

	add()	
	ave()

}

func add(){

var a,b int

fmt.Scanf("%d,%d",&a,&b)

c :=a+b

fmt.Printf("the addition is %d",c)

}
func ave(){

var a,b,c int

fmt.Printf("Enter three numbers\n")

fmt.Scanf("%d,%d,%d",&a,&b,&c)

ave :=(a+b+c)/3

fmt.Printf("Average value is %d",ave)

}
