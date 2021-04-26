package main

import "fmt"

type Node struct {
	data string
	left *Node
	right *Node
}

/*
广度优先遍历(层序遍历)
结果存在result里面,如果不存在可以少一层变量
*/
func breadthFirstSearch(node Node) []string{
	var result []string
	var nodes []Node = []Node{node}
	fmt.Println(len(nodes))
	for len(nodes) > 0{
		node := nodes[0]
		fmt.Println(nodes)
		result = append(result,node.data)
		if (node.left != nil){
			nodes = append(nodes,*node.left)
		}
		if (node.right != nil){
			nodes = append(nodes,*node.right)
		}
	}
	return result
}

func preOrderRecursive(n int,node Node){

	if n != 1 && n != 2 && n != 3{
		fmt.Println("参数错误")
		return
	}

	//前序
	if n == 1{
		fmt.Println(node.data)
	}
	if node.left != nil{
		preOrderRecursive(n,*node.left)
	}

	//中序
	if n == 2{
		fmt.Println(node.data)
	}
	if node.right != nil{
		preOrderRecursive(n,*node.right)
	}

	//后序
	if n == 3{
		fmt.Println(node.data)
	}
}


/*
广度优先: a -> b -> c -> d -> f -> e -> g
先序遍历: a -> b -> d -> e -> f -> g -> c
中序遍历: e -> d -> b -> g -> f -> a -> c
后序遍历: e -> d -> g -> f -> b -> c -> a
*/
func main() {
	nodeG := Node{"g",nil,nil}
	nodeF := Node{"f",&nodeG,nil}
	nodeE := Node{"e",nil,nil}
	nodeD := Node{"d",&nodeE,nil}
	nodeC := Node{"c",nil,nil}
	nodeB := Node{"b",&nodeD,&nodeF}
	nodeA := Node{"a",&nodeB,&nodeC}

	//fmt.Println("广度优先:",breadthFirstSearch(nodeA))

	//fmt.Println("递归前序遍历:")
	//preOrderRecursive(1,nodeA)
	//
	//fmt.Println("递归中序遍历:")
	//preOrderRecursive(2,nodeA)
	//
	//fmt.Println("递归后序遍历:")
	//preOrderRecursive(3,nodeA)

	fmt.Println("非递归前序遍历:",preOrderLoop(&nodeA))
	fmt.Println("非递归中序遍历:",midOrderLoop(&nodeA))
	fmt.Println("非递归后序遍历:",postOrderLoop(&nodeA))
}

type seqStack struct {
	data [100]*Node
	tag [100]int //后续遍历准备
	top int //数组下标
}

func preOrderLoop(node *Node)(result []string){
	var s seqStack
	s.top = -1 //空
	if node == nil{
		panic("no data here")
	}else{
		for node != nil || s.top != -1{
			for node != nil{
				result = append(result,node.data)
				s.top++
				s.data[s.top] = node
				node = node.left
			}
			s.top--
			node = s.data[s.top + 1]
			node = node.right
		}
	}
	return
}

func midOrderLoop(node *Node) (result []string) {
	var s seqStack
	s.top = -1
	if node == nil {
		panic("no data here")
	}else {
		for node != nil || s.top != -1 {
			for node != nil {
				s.top++
				s.data[s.top] = node
				node = node.left
			}
			s.top--
			node = s.data[s.top + 1]
			result = append(result, node.data)
			node = node.right
		}
	}
	return
}




func postOrderLoop(node *Node) (result []string) {
	var s seqStack
	s.top = -1

	if node == nil {
		panic("no data here")
	} else {
		for node != nil || s.top != -1 {
			for node != nil {
				s.top++
				s.data[s.top] = node
				s.tag[s.top] = 0
				node = node.left
			}

			if s.tag[s.top] == 0 {
				node = s.data[s.top]
				s.tag[s.top] = 1
				node = node.right
			} else {
				for s.tag[s.top] == 1 {
					s.top--
					node = s.data[s.top+1]
					result = append(result, node.data)
					if s.top < 0 {
						break
					}
				}
				node = nil
			}
		}
	}
	return
}




