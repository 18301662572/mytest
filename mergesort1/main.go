package main

import (
	"fmt"
	"math/rand"
)

//归并排序

//创建一个16位数的切片，并对其进行对并排序

func mergeSort(data []int)[]int{
	if len(data)<=1{
		return data
	}
	num:=len(data)/2
	left:=mergeSort(data[:num])
	right:=mergeSort(data[num:])
	return merge(left,right)
}

func merge(left,right []int)(result[]int){
	l,r:=0,0
	for len(left)>l && len(right)>r{
		if left[l]<right[r]{
			result=append(result,left[l])
			l++
		}else{
			result=append(result,right[r])
			r++
		}
	}
	result=append(result,left[l:]...)
	result=append(result,right[r:]...)
	return
}

func main(){
	aa:=make([]int,0,16)
	for i:=0;i<16 ;i++  {
		aa=append(aa,rand.Intn(100))
	}
	fmt.Println(aa)
	data:=mergeSort(aa)
	fmt.Println(data)
}
