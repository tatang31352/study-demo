package main

import (
	"fmt"
)

type Node struct {
	Data int
	Prev *Node
	Next *Node
}


type DeLinkList struct {
	head *Node //由于链表任何操作都需要从头部开始，因此这里定义一个链表头部属性
}

//只需要判断首部节点是否为空
func (self *DeLinkList) IsEmpty() bool{
	return self.head == nil
}

//链表长度
func (self *DeLinkList) Length() int{
	count := 0
	head := self.head
	for head != nil{
		count += 1
		head = head.Next
	}
	return count
}

//从头遍历
func (self *DeLinkList) Range()  {
	head := self.head
	for head != nil{
		fmt.Println("当前元素",head.Data)
		head = head.Next
	}
}

//从头部添加节点
func (self *DeLinkList) Add(value int){
	newNode := &Node{Data:value}
	if self.head == nil{
		self.head = newNode
	}else {
		newNode.Next = self.head
		self.head.Prev = newNode
		self.head = newNode
	}
}

//从尾部添加元素
func (self *DeLinkList) Append(value int){
	head := self.head
	for head.Next != nil{
		head = head.Next
	}
	newNode := &Node{Data:value}
	newNode.Prev = head
	head.Next = newNode
}

//删除任意位置的元素，todo 需要处理：删除的元素在头部，中部，尾部三种情况
func (self *DeLinkList) Delete(value int){
	head := self.head
	//删除头部
	if head.Data == value{
		self.head = head.Next
	}else {
		for head != nil{
			if head.Data == value{
				//修改这个元素之前的一个元素Next指针
				head.Prev.Next = head.Next
				if head.Next != nil{
					head.Next.Prev = head.Prev
				}
				break
			}else {
				head = head.Next
			}
		}

	}
}

// 在任意位置插入数据 todo 需要考虑插入数据位于 首部, 尾部,中间的处理情况
func (self *DeLinkList) Insert(index,value int){
	if index <= 0{
		self.Add(value)
	}else if index > self.Length(){
		self.Append(value)
	}else{
		//插入数据在中间的情况,需要找到index之前的元素修改指针
		head := self.head
		count := 0
		for count < index{
			count += 1
			head = head.Next
		}
		newNode := &Node{Data:value}
		newNode.Prev = head.Prev
		newNode.Next = head
		head.Prev.Next = newNode
		head.Prev = newNode
	}
}

func main() {
	delink := &DeLinkList{}
	delink.Add(10)
	delink.Add(15)
	delink.Append(8)
	delink.Append(66)
	delink.Append(15)
	delink.Insert(2,88)
	fmt.Println("长度",delink.Length())
	delink.Range()
}
