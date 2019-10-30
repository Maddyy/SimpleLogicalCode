package main

import (
	"fmt"
	"os"
	"strconv"
)

func PrimeNumber(){
	flag := false
  num1 := os.Args[1]
  num,_ :=strconv.Atoi(num1)
 var val int=(num/2)
  for i :=2; i<=val; i++{ 
  if(num%i == 0){
	fmt.Println("number is not prime")
	flag=true
	break
	  }
	  
  }
  if(flag==false){
	fmt.Println("number is  prime")
  }
}