package main
import (
	"fmt"
	"os"
	"strconv"
)
func factorial(){
	var fact int=1
  input :=os.Args[1]
  num,_ :=strconv.Atoi(input)
  for i:=1;i<=num;i++{
  fact=fact*i
  }
  fmt.Println("factorial ",fact)
}
