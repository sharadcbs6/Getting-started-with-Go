package main
import "fmt"
type Author struct{
	name string
	branch string
	language string
}

func main(){
	a1 :=  Author { name: "Harish" , branch: "CSE" , language: "Golang" };
	a2 := Author { name: "Harish" , branch: "CSE" , language: "Golang" };
	a3 := Author { name:"Shivam" , branch: "CSE" , language: "Golang" };
	if a1==a2 || a2==a3 {
		fmt.Println("Both authors are same")
	} else {
		fmt.Println("Both authors are different"); 
		/* learnings -> else block start from the closing brace */
	}
	 
}
