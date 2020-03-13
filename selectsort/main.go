package main

import (
	"fmt"
	"math/rand"
)

//选择排序

func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main(){
	aa:=make([]int,0,6)
	for i:=0;i<6 ;i++  {
		aa=append(aa,rand.Intn(100))
	}
	fmt.Println(aa)
	selectionSort(aa)
	fmt.Println(aa)
}
