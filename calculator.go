package main
import "fmt"
func main(){
    subtraction()
	}
    func subtraction(){
	var a,b int
	
	fmt.Printf("enter two numbers\n")
	fmt.Scanf ("%d,%d",&a,&b)
	sub:=a-b
	fmt.Printf("the result is %d",sub)
	}
}
