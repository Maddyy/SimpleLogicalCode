package main

import (
	"fmt"
	
)
func main(){
	input :="Hello Sharad How are you"
	output :=Reverse(input)
	fmt.Println(output)
	
}
func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
  }
