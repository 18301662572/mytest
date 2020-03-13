package main

import (
	"fmt"
	"math/rand"
)

//快速排序
//创建一个10个数字的正整形切片，对其进行快速排序

func quickSort(data []int){
	if len(data)<=1{
		return
	}
	base:=data[0]
	r:=len(data)-1 //从后往前的下标
	l:=0 //base基数从前往后的下标
	for i:=1;i<=r ;  {
		if data[i]>base{
			data[i],data[r]=data[r],data[i]
			r--
		}else{
			data[i],data[l]=data[l],data[i]
			l++
			i++
		}
	}
	quickSort(data[:l])
	quickSort(data[l+1:])
}

func main(){
	aa:=make([]int,0,10)
	for i:=0;i<10 ;i++  {
		aa=append(aa,rand.Intn(100))
	}
	fmt.Println(aa)
	quickSort(aa)
	fmt.Println(aa)
}
