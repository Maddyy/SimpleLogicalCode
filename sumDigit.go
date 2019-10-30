package main

import (
	"fmt"
	"os"
	"strconv"
)
func sumDigit(){
	var sum int=0
	input :=os.Args[1]
	num,_ :=strconv.Atoi(input)
	
	for num >0{
		r :=num%10
		sum= sum+r
		num=num/10
	}
	fmt.Println("Sum is ",sum)
}
