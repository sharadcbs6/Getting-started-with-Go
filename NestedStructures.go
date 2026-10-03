package main
import "fmt"

type Address struct{
	Street  string
	City 	string
	State 	string
	PostalCode string
}

type Person struct{
	FirstName string
	LastName string
	Age		 int 
	Address  Address
}
func main(){
	p :=  Person{
		FirstName : "Sharad",
		LastName : "Yadav",
		Age 	 : 22,
		Address  : Address{
			Street : " XY ",
			City   : "New Delhi",
			State  : "Delhi",
			PostalCode : "110061",
		},
	};
	fmt.Println("Name:",p.FirstName,p.LastName);
	fmt.Println("Age:",p.Age);
	fmt.Println("Address:",p.Address.Street,",",p.Address.City,",",p.Address.State,",",p.Address.PostalCode);
 }