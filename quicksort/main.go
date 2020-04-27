package main

import (
	"fmt"
)

//快速排序

var count=0

func quickSort(data []int) {
	c1:=0
	if len(data) <= 1 {
		return
	}
	base := data[0]
	l, r := 0, len(data)-1
	for i := 1; i <= r; {
		count++
		c1++
		if data[i] > base {
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			data[i], data[l] = data[l], data[i]
			l++
			i++
		}
	}
	fmt.Println(data)
	fmt.Println("func c1:",c1)
	//fmt.Println("func coount:",count)
	quickSort(data[:l])
	quickSort(data[l+1:])
}

func main() {
	//s := make([]int, 0, 8)
	//for i := 0; i < 8; i++ {
	//	s = append(s, rand.Intn(100))
	//}
	//fmt.Println(s)
	//s:=[]int{8,7,6,5,4}    //10
	//s:=[]int{4,5,6,7,8}    //8
	s:=[]int{7,4,6,8,5}		 //7
	quickSort(s)
	fmt.Println(s)
	fmt.Println("main coount:",count)
}
