package main


import (
	"fmt"
	"os"
	"strconv"
)
func Palindorme(){
input :=os.Args[1]
num,_  :=strconv.Atoi(input)
tmp :=num
var sum int
for num>0 {
	r :=num%10
	sum =sum*10+r
	num =num/10
}
if(sum==tmp){
fmt.Println("Palindrom")
} else{
	fmt.Println("NOT Palindrom")
}
}