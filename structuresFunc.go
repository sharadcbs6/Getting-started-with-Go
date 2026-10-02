package main
import "fmt"
type Address struct{
	Name string
	City string
	Pincode int

}

func main(){

	a1:=Address{"Sharad","Delhi",110061};
	a2:=Address{"Anshul","Gurgaon",122017};
	fmt.Println(a1);
	fmt.Println(a2);
}