package greetings
import (
	"errors"
	"fmt"

)
func Hello(name string) (string,error){
	//if no name was given ,return a new error message
	if name==""{
		return "",errors.New("empty Name");
	}
	message:=fmt.Sprintf("Hello, %v. Weclome! ",name);
	return message,nil;
}