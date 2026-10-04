package main
import "fmt"
type Person struct{
	Name string
	Greet func() string
}
func main(){
	p := Person{ Name: "Sharad Yadav",};
	p.Greet= func () string{
		return "Hello, "+p.Name;
	} 
	fmt.Println(p.Greet());

}