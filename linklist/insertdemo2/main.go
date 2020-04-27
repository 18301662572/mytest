package main

import "fmt"

//插入节点
//单链表的节点插入方法一般使用头插法或者尾插法。
//尾插法：每次插入节点在尾部，这也是我们较为习惯的方法

type Node struct {
	date int
	next *Node
}

//遍历链表，循环输出
func Shownode(node *Node){
	for node!=nil{
		fmt.Println(*node)
		node=node.next //移动指针
	}
}

func main() {
	head:=new(Node)
	head.date=0
	var tail *Node
	tail=head  //tail用于记录最末尾的节点的地址，刚开始tail的的指针指向头节点
	for i:=1;i<=10 ;i++  {
		var node=Node{date:i}
		tail.next=&node
		tail=&node
	}
	Shownode(head) //遍历结果
}

