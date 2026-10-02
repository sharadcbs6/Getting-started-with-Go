package main
import "fmt"
func main(){
	p:=34;
	q:=32;
	result1:=p&q //bitwise and
	fmt.Printf("Result of p & &q = %d \n",result1);
	result2:=p|q //bitwise or
	fmt.Printf("Result of p | q = %d \n",result2);
	result3:=p^q //bitwise xor
	fmt.Printf("Result of p ^ q = %d \n",result3);
	result4:=p<<1 //left shift
	fmt.Printf("Result of p << 1 = %d \n",result4);
	result5:=p>>1 //right shift
	fmt.Printf("Result of p >> 1 = %d \n",result5);
	result6:=p&^q //and not
	fmt.Printf("Result of p &^ q = %d ",result6);
}