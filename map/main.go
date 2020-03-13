package main

import "sort"

//map for range

func main(){
	var m=map[string]int{
		"hello":0,
		"morning":1,
		"my":2,
		"girl":3,
	}
	var keys []string
	for key := range m {
		keys=append(keys,key)
	}
	sort.Strings(keys)
	for _,k := range keys {
		println("Key:",k,"Value:",m[k])
	}
}
