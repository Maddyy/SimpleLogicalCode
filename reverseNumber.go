package main 
import (
	"fmt"
	"os"
	"strconv"
)
func reverNum(){
	var newNum int=0
	input :=os.Args[1]
	num,_   :=strconv.Atoi(input)
   	for num>0{
	rev :=num%10
	newNum =newNum*10+rev
	num=num/10
   }
   fmt.Println("new",newNum)
}
