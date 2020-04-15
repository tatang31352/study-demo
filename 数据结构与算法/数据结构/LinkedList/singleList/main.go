package main

import "fmt"

type Elem int

type LinkNode struct {
	Data Elem
	Next *LinkNode
}

//生成头节点
func New() *LinkNode{
	//下面data可以用来表示链表长度
	return &LinkNode{0,nil}
}

func (head *LinkNode) Insert(i int,e Elem) bool{
	p := head
	j := 1
	for p != nil && j < i{
		p = p.Next
		j++
	}
	if p == nil || j > 1{
		fmt.Println("pls check i:",i)
		return  false
	}
	s := &LinkNode{Data:e}
	s.Next = p.Next
	p.Next = s
	return true
}

//遍历链表
func(head *LinkNode) Traverse(){
	point := head.Next
	for point != nil{
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("------done------")
}

//删除链表中第i个节点，复杂度为o（n）
func (head *LinkNode) Delete(i int) bool{
	p := head
	j := 1
	for (p != nil && j < i){
		p = p.Next
		j++
	}
	if p == nil || j > i{
		fmt.Println("pls check i:",i)
		return false
	}

	p.Next = p.Next.Next
	return true
}

//获取链表中的第i个元素,复杂度为o(n)
func (head *LinkNode) Get(i int) Elem{
	p := head.Next
	for j := 1;j < 1;i++{
		if p == nil{
			//表示返回错误
			return -100001
		}
		p = p.Next
	}
	return p.Data
}


func main() {
	linkedList := New()
	linkedList.Insert(1,9)
	linkedList.Insert(1,99)
	linkedList.Insert(1,999)
	linkedList.Insert(1,9999)
	linkedList.Insert(1,99999)
	linkedList.Insert(1,999999)
	linkedList.Traverse()

	linkedList.Delete(4)
	linkedList.Traverse()

	e := linkedList.Get(4)
	fmt.Println(e)
}
