package main

import (
	"fmt"
	"os"
)

type Emp struct {
	ID int
	Name string
	Next *Emp
}

//第一个节点就存放员工
type EmpLink struct {
	Head *Emp
}

//定义HashTable
type HashTable struct {
	LinkArr [7]EmpLink
}

//添加员工的方法
func (empl *EmpLink) InsertEmp(emp *Emp){
	cur := empl.Head
	var pre *Emp = nil
	if cur == nil{
		empl.Head = emp
		return
	}
	//如果不是一个空链表,找到对应的位置并插入
	for {
		if cur != nil{
			if cur.ID >= emp.ID{
				break
			}
			pre = cur
			cur = cur.Next
		}else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}

func (hash *HashTable) Insert(emp *Emp){
	//使用散列函数,确定将员工添加到哪个链表
	linkNum := hash.HashFunc(emp.ID)
	hash.LinkArr[linkNum].InsertEmp(emp)
}

func (empl *EmpLink) FindByID(id int) *Emp {
	cur := empl.Head
	for {
		if cur != nil && cur.ID == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

func (hash *HashTable) Find(id int) *Emp{
	//使用散列函数确定在哪个链表
	linkNum := hash.HashFunc(id)
	return hash.LinkArr[linkNum].FindByID(id)
}

func (hash *HashTable) HashFunc(id int) int{
	return id % 7
}

func (empl *EmpLink) ShowLink(num int){
	if empl.Head == nil{
		fmt.Printf("当前%d链表为空\n",num)
		return
	}
	//否则遍历显示数据
	cur := empl.Head
	for{
		if cur != nil{
			fmt.Printf("链表:%d,员工id:%d,员工名字:%s-->",num,cur.ID,cur.Name)
			cur = cur.Next
		}else{
			break
		}
	}
	fmt.Println(``)
}

func (hash *HashTable) Show(){
	for i := 0; i < len(hash.LinkArr);i++{
		hash.LinkArr[i].ShowLink(i)
	}
}

func (emp *Emp) ShowMe(){
	fmt.Printf("链表%d找到该员工%d\n",emp.ID%7,emp.ID)
}



func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for{
		fmt.Println("==========员工菜单==========")
		fmt.Println("insert 表示添加员工")
		fmt.Println("show 表示显示员工")
		fmt.Println("find 表示查询员工")
		fmt.Println("exit 表示退出员工")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "insert":
			fmt.Println("请输入员工id:")
			fmt.Scanln(&id)
			fmt.Println("请输入员工名字:")
			fmt.Scanln(&name)
			emp := &Emp{
				ID:id,
				Name:name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.Show()
		case "find":
			fmt.Println("请输入你要查找的id:")
			fmt.Scanln(&id)
			emp := hashTable.Find(id)
			if emp == nil{
				fmt.Printf("id=%d的员工不存在",id)
			}else{
				//显示雇主信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		}
	}

}
