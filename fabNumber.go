package main

import (
	"fmt"
	"os"
	"strconv"
)

func mainx() {
	 var a,b int = 0,1
	 var c int
	 febNumber := os.Args[1]
	 num,_ :=strconv.Atoi(febNumber)
	 for i :=2;i<num; i++ {
        c=a+b
		fmt.Print(" ",c)
        a=b
	    b=c
	}
}
