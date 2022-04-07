package main

import (
	"fmt"

)



func main() {

	
	var a,b int= 1,2

	sum(a,b)

}

func sum(a,b int)  int{

	var c  int
	c = a+b
	fmt.Printf("leijie = %d",c)
	return c

}


