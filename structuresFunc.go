package main
import "fmt"
// type Address struct{
// 	Name string
// 	City string
// 	Pincode int

// }



// func main(){

// 	a1:=Address{"Sharad","Delhi",110061};
// 	a2:=Address{"Anshul","Gurgaon",122017};
// 	fmt.Println(a1);
// 	fmt.Println(a2);
// }


 type Car struct {
	Name, Model, Color string
	WeightInKg	float64
}


func main(){
	c := Car { Name: "Mustang",  Model : "Mustang GT" , Color : "Metallic Grey", WeightInKg : 190.10} ;
	fmt.Println("Name: "+ c.Name);
	fmt.Println("Model: "+ c.Model);
	fmt.Println("Color: "+c.Color);
	fmt.Println("Weight: ",c.WeightInKg);
}