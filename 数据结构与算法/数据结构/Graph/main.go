package main

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

type Graph struct {
	nodes  []*Node   //节点集
	edges map[Node][]*Node //邻接表表示的无向图
	lock sync.RWMutex //保证线程安全
}

//增加节点
func (g *Graph) AddNode(n *Node){
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes,n)
}

//增加边
func (g *Graph) AddEdge(u,v *Node){
	g.lock.Lock()
	defer g.lock.Unlock()
	//首次建立图
	if g.edges == nil{
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*u] = append(g.edges[*u],v)//建立u->v的边
	g.edges[*v] = append(g.edges[*v],u)//由于是无向图，同时存在v->u的边
}

//输出图
func (g *Graph) String(){
	g.lock.RLock()
	defer g.lock.RUnlock()
	str := ""
	for _, iNode := range g.nodes{
		str += iNode.String() + " -> "
		nexts := g.edges[*iNode]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}


//输出节点
func (n *Node) String() string{
	return fmt.Sprintf("%v",n.value)
}


func main() {
	g := Graph{}
	n1,n2,n3,n4,n5 := Node{1},Node{2},Node{3},Node{4},Node{5}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1,&n2)
	g.AddEdge(&n1,&n5)
	g.AddEdge(&n2,&n3)
	g.AddEdge(&n2,&n4)
	g.AddEdge(&n2,&n5)
	g.AddEdge(&n3,&n4)
	g.AddEdge(&n4,&n5)

	g.String()

	g.BFS(func(node *Node) {
		fmt.Printf("[Current Traverse Node]: %v\n",node)
	})
}


type NodeQueue struct {
	nodes []Node
	lock  sync.RWMutex
}

// 实现 BFS 遍历
func (g *Graph) BFS(f func(node *Node)) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	// 初始化队列
	q := NewNodeQueue()
	// 取图的第一个节点入队列
	head := g.nodes[0]
	q.Enqueue(*head)
	// 标识节点是否已经被访问过
	visited := make(map[*Node]bool)
	visited[head] = true
	// 遍历所有节点直到队列为空
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		nexts := g.edges[*node]
		// 将所有未访问过的邻接节点入队列
		for _, next := range nexts {
			// 如果节点已被访问过
			if visited[next] {
				continue
			}
			q.Enqueue(*next)
			visited[next] = true
		}
		// 对每个正在遍历的节点执行回调
		if f != nil {
			f(node)
		}
	}
}

// 生成节点队列
func NewNodeQueue() *NodeQueue {
	q := NodeQueue{}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = []Node{}
	return &q
}

// 入队列
func (q *NodeQueue) Enqueue(n Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = append(q.nodes, n)
}

// 出队列
func (q *NodeQueue) Dequeue() *Node {
	q.lock.Lock()
	defer q.lock.Unlock()
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return &node
}

// 判空
func (q *NodeQueue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.nodes) == 0
}