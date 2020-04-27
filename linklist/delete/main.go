package main

import "errors"

//删除节点

type ListNode struct {
	data    int
	next *ListNode

}

func (l *ListNode)DelNode(d int) {
	if l==nil{
		errors.New("Empty List!")
		return
	}
	for l.next!=nil{
		if l.next.data==d{
			l.next=l.next.next
			//return  此处控制找到相同数据是否全部删除操作
		}
		l=l.next
	}
}

//基于位置删除元素
func (head *ListNode)DelNodeById(id int) {
	i:=1
	node:=head.next
	tmp:=head
	for node!=nil {
		if i==id{
			tmp.next=node.next
		}
		i++
		tmp=node
		node=node.next
	}
	if i<id{
		errors.New("the len of the list is less !")
	}

}



