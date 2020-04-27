package main

import "fmt"

//链表demo
//单链表：每个节点包含下一个节点的地址，这样把所有节点都串起来的链式数据数据结构叫做链表，
// 通常把链表中的第一个节点叫做表头。

//使用struct定义单链表
type Node struct {
	date int
	next *Node
}

//遍历链表，循环输出
//链表的遍历是通过移动指针进行遍历，当指针到最后一个节点时，其next指针为nil
func Shownode(node *Node){
	for node!=nil{
		fmt.Println(*node)
		node=node.next //移动指针
	}
}


func main() {
	head:=new(Node)
	head.date=0

	node1:=new(Node)
	node1.date=1
	head.next=node1

	node2:=new(Node)
	node2.date=2
	node1.next=node2

	//遍历链表，循环输出
	Shownode(head)
}
