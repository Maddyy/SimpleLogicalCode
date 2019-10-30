package main

import (
	"fmt"
	"os"
	"strconv"
)
func factRecu(){
	input :=os.Args[1]
	num , _ :=strconv.Atoi(input)
   factval := Factorial(num)
   fmt.Println("Fact",factval)
}
func Factorial(num1 int)(number int){
 if(num1==0){
	 return 1
 }else {
	return  num1*Factorial(num1-1)
	}
}

