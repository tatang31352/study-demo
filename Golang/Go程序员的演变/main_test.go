package main

import "testing"


// 初级程序员
func TestFactorial(t *testing.T) {
	cases := []struct{
		n int
		res int
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := Factorial(c.n)
		if result != c.res{
			t.Fatalf("Factorial function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}

// 函数式程序员
func TestFunction(t *testing.T) {
	cases := []struct{
		n int
		res int
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := Function(c.n)
		if result != c.res{
			t.Fatalf("Funtion function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}

// 泛型Go程序员
func TestExtensive(t *testing.T) {
	cases := []struct{
		n interface{}
		res interface{}
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := Extensive(c.n)
		if result != c.res{
			t.Fatalf("Extension function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}

// 多线程Go程序员
func TestMulti(t *testing.T) {
	cases := []struct{
		n int
		res int
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := Multi(c.n)
		if result != c.res{
			t.Fatalf("Multi function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}

// 发现型Go模式暂未测试
/*func TestFind(t *testing.T) {
	cases := []struct{
		n int
		res <-chan
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := Find(c.n)
		if result != c.res{
			t.Fatalf("Extension function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}*/

// 使用成熟的解决方案修复Go缺陷
func TestCalculateFactorial(t *testing.T) {
	cases := []struct{
		n int
		res int
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := NewFactorial(c.n).CalculateFactorial()
		if result != c.res{
			t.Fatalf("CalculateFactorial function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}

// 高级Go程序员
func TestSenior(t *testing.T) {
	cases := []struct{
		n int
		res int
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := Senior(c.n)
		if result != c.res{
			t.Fatalf("Senior function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}

// Rob Pike
func TestRobPike(t *testing.T) {
	cases := []struct{
		n int
		res int
	}{
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
	}

	for _,c:= range cases{
		result := RobPike(c.n)
		if result != c.res{
			t.Fatalf("RobPike function failed,n: %d,res: %d,result:%d",c.n,c.res,result)
		}
	}
}