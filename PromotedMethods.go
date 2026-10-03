package main
import "fmt"

type details struct{
	name string
	age  int 
	gender string
	psalary int
}

type employee struct{
	post string
	eid  int 
	details
}
func (d details)  promotedmethod(tsalary int) int{
	return d.psalary * tsalary;
}
func main(){
	values:= employee{
		post: "SAP Consultant",
		eid: 10842131,
		details: details{
			name: "Sharad Yadav",
			age: 22,
			gender: "Male",
			psalary: 935,	
		},
	}
	fmt.Println("Name: ",values.name);
	fmt.Println("Age:",values.age);
	fmt.Println("Gender:",values.gender);
	fmt.Println("Per Day Salary:",values.psalary);
	fmt.Println("Montly In Hand Salary:",values.promotedmethod(30));
}