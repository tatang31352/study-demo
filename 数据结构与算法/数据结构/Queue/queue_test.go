package Queue

import (
	"errors"
	"reflect"
	"testing"
)


func TestQueue_Push(t *testing.T) {
	data  := []*Queue{
		{2,3,4},
		{"a","b","c"},
		{2,"2",9},
	}

	ress := []*Queue{
		{2,3,4,8},
		{"a","b","c",8},
		{2,"2",9,8},
	}

	cases := []struct{
		q *Queue
		x interface{}
		res *Queue
	}{
		{data[0],8,ress[0]},
		{data[1],8,ress[1]},
		{data[2],8,ress[2]},
	}

	for _,c := range cases{
		c.q.Push(c.x)
		if !reflect.DeepEqual(c.res,c.q) {
			t.Fatalf("Push function failed,q:%d,x:%d,res:%d",c.q,c.x,c.res)
		}
	}
}

func TestQueue_Pop(t *testing.T) {
	var data  = []*Queue{
		{},
		{"a","b","c"},
		{2,"2",9},
	}

	cases := []struct{
		q *Queue
		top interface{}
		err error
	}{
		{data[0],nil,errors.New("Can`t pop an empty queue")},
		{data[1],"a",nil},
		{data[2],2,nil},
	}

	for _,c := range cases{
		c1,_ := c.q.Pop()
		if c1 != c.top  {
			t.Fatalf("Pop function failed,q:%d,x:%d,res:%d",c.q,c.top,c.err)
		}
	}

}

func TestQueue_Top(t *testing.T) {
	var data  = []*Queue{
		{2,3,4},
		{"a","b","c"},
		{2,"2",9},
	}

	cases := []struct{
		q *Queue
		top interface{}
		err error
	}{
		{data[0],2,nil},
		{data[1],"a",nil},
		{data[2],2,nil},
	}

	for _,c := range cases{
		c1,_ := c.q.Pop()
		if c1 != c.top  {
			t.Fatalf("Top function failed,q:%d,x:%d,res:%d",c.q,c.top,c.err)
		}
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	var data  = []Queue{
		{2,3,4},
		{"a","b","c"},
		{},
	}

	cases := []struct{
		q Queue
		err bool
	}{
		{data[0],false},
		{data[1],false},
		{data[2],true},
	}

	for _,c := range cases{
		if c.q.IsEmpty() != c.err   {
			t.Fatalf("Pop function failed,q:%d,err:%v",c.q,c.err)
		}
	}
}

func TestQueue_Len(t *testing.T){
	var data  = []Queue{
		{2,3,4},
		{"a","b","c",4},
		{2,"2",9},
	}

	cases := []struct{
		q Queue
		n int
	}{
		{data[0],3},
		{data[1],4},
		{data[2],3},
	}

	for _,c := range cases{
		if c.q.len() != c.n  {
			t.Fatalf("len function failed,q:%d,n:%d",c.q,c.n)
		}
	}
}