package main

import "fmt"

//defer demo

func main(){
	a:=1
	defer fmt.Println("the value of a1:",a)
	a++
	defer func(){
		fmt.Println("the value of a2:",a)
	}()
}


