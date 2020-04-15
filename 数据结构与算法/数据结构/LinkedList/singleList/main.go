package main

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

func main() {
	
}
