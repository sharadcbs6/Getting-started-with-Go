package main
import "fmt"
type details struct{
	name   string
	age    int
	gender string
}
type student struct{
	branch string
	year   int
	details
}
func main(){
	values:= student{
		branch: "CS",
		year : 2021,
		details: details{
			name: "Sharad",
			age: 17,
			gender: "Male",
		},
	}

	fmt.Println("Name:",values.name);
	fmt.Println("Age:",values.age);
	fmt.Println("Gender:",values.gender);
	fmt.Println("Branch:",values.branch);
	fmt.Println("Year:",values.year);
}