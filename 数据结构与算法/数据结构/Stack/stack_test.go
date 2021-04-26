package Stack

import (
	"reflect"
	"testing"
)

func TestStack_Len(t *testing.T){

	cases := []struct{
		s Stack
		i int
	}{
		{Stack{1,2,3},3},
		{Stack{1,2,3,"2"},4},
		{Stack{1,2,3},3},
		{Stack{1,2,3},3},
	}

	for _,c := range cases{
		if c.s.len() != c.i {
			t.Fatalf("lem function failed,s:%d,c:%d",c.s,c.i)
		}
	}
}

func TestStack_IsEmpty(t *testing.T) {
	cases := []struct{
		s Stack
		ok bool
	}{
		{Stack{1,2,3,},false},
		{Stack{1,2,3,},false},
		{Stack{1,2,3,},false},
		{Stack{},true},
	}

	for _,c := range cases{
		if c.s.IsEmpty() != c.ok{
			t.Fatalf("IsEmpty function failed,s:%d,ok:%v",c.s,c.ok)
		}
	}
}

func TestStack_Cap(t *testing.T) {
	cases := []struct{
		s Stack
		n int
	}{
		{Stack{2,3,4,},3},
		{Stack{2,3,4,},3},
		{Stack{2,3,},2},
		{Stack{},0},
	}

	for _,c := range cases{
		if c.s.Cap() != c.n{
			t.Fatalf("Cap function failed,s:%d,n:%d",c.s,c.n)
		}
	}
}

func TestStack_Push(t *testing.T) {
	cases := []struct{
		s *Stack
		v interface{}
		r *Stack
	}{
		{&Stack{1,2,3,},3,&Stack{1,2,3,3,}},
		{&Stack{1,2,3,3,},4,&Stack{1,2,3,3,4,}},
		{&Stack{1,2,3,5,},4,&Stack{1,2,3,5,4}},
		{&Stack{1,2,3,"sssssss",},4,&Stack{1,2,3,"sssssss",4}},
	}

	for _,c := range cases{
		 c.s.Push(c.v)
		if !reflect.DeepEqual(*c.s,*c.r){
			t.Fatalf("push function failed,s:%d,v:%d,r:%d",c.s,c.v,c.r)
		}
	}
}

func TestStack_Top(t *testing.T) {
	cases := []struct{
		s Stack
		v interface{}
	}{
		{Stack{1,2,3,},3},
		{Stack{1,2,4,},4},
		{Stack{1,2,5,},5},
		{Stack{1,2,6,},6},
		{Stack{},nil},
	}

	for _,c := range cases{
		res,_ := c.s.Top()
		if c.v != res{
			t.Fatalf("Top function failed,s:%d,v:%d",c.s,c.v)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	cases := []struct{
		s *Stack
		v interface{}
		res *Stack
	}{
		{&Stack{1,2,3,},3,&Stack{1,2,}},
		{&Stack{1,2,3,},3,&Stack{1,2,}},
		{&Stack{1,2,3,},3,&Stack{1,2,}},
		{&Stack{},nil,&Stack{}},
	}

	for _,c := range cases{
		res,_ := c.s.Pop()
		if c.v != res || !reflect.DeepEqual(c.s,c.res){
			t.Fatalf("Pop function failed,s:%d,v:%d",c.s,c.v)
		}
	}
}

