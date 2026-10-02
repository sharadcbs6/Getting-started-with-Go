package main
import "fmt"
func main(){
	value:=func(){
		fmt.Println("Hello");
	};
	func(ele string){
		fmt.Println(ele);
	}("Hello Sharad ");
}