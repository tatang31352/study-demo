package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

const n int = 0

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

const (
	col = 10000
	row = 10000
)

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {  //监控cpu
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}


	//初级Go程序员
	bT := time.Now()
	res := Factorial(n)
	eT := time.Since(bT)
	fmt.Printf("初级Go程序员计算%v的阶乘结果为:%v,共耗时:%v\n",n,res,eT)


	//函数式Go程序员
	bT = time.Now()
	res = Function(n)
	eT = time.Since(bT)
	fmt.Printf("函数式Go程序员计算%v的阶乘结果为:%v,共耗时:%v\n",n,res,eT)

	//泛型Go程序员
	bT = time.Now()
	resI := Extensive(n)
	eT = time.Since(bT)
	fmt.Printf("泛型Go程序员计算%v的阶乘结果为:%v,共耗时:%v\n",n,resI,eT)

	//多线程优化的Go程序员
	bT = time.Now()
	res = Multi(n)
	eT = time.Since(bT)
	fmt.Printf("多线程优化的Go程序员计算%v的阶乘结果为:%v,共耗时:%v\n",n,res,eT)

	//发现型Go模式
	bT = time.Now()
	ch := Find(n)
	for i:= 0;i < n;i++{
		if i == n - 1{
			fmt.Printf("发现型Go模式计算%v的阶乘结果为:%v,共耗时:%v\n",n,<-ch,eT)
		}else{
			<-ch
		}
	}


	//使用成熟的解决方案修复Go缺陷
	bT = time.Now()
	resN := NewFactorial(n)
	res = resN.CalculateFactorial()
	eT = time.Since(bT)
	fmt.Printf("使用成熟的解决方案修复Go缺陷计算%v的阶乘结果为:%v,共耗时:%v\n",n,res,eT)

	//高级Go程序员
	bT = time.Now()
	res = Senior(n)
	eT = time.Since(bT)
	fmt.Printf("高级程序员计算%v的阶乘结果为:%v,共耗时:%v\n",n,res,eT)

	//Rob Pike
	bT = time.Now()
	res = Senior(n)
	eT = time.Since(bT)
	fmt.Printf("Rob Pike计算%v的阶乘结果为:%v,共耗时:%v\n",n,res,eT)


	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // GC，获取最新的数据信息
		if err := pprof.WriteHeapProfile(f); err != nil {  // 写入内存信息
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

// 初级Go程序员
func Factorial(n int) int{
	res := 1

	for i := 1; i <= n;i++{
		res *= i
	}

	return res
}


//函数式Go程序员
func Function(n int) int{
	if n == 0{
		return 1
	}else{
		return Function(n - 1) * n
	}
}

//泛型Go程序员
func Extensive(n interface{}) interface{}{
	v,valid := n.(int)
	if !valid {
		return 0
	}
	res := 1

	for i := 1;i <= v;i++{
		res *= i
	}

	return res
}

//多线程优化的Go程序员
func Multi(n int) int{
	if n == 1{
		return 1
	}

	var(
		left,right = 1,1
		wg sync.WaitGroup
	)
	wg.Add(2)

	pivot := n / 2

	go func() {
		for i := 1;i < pivot;i++{
			left *= i
		}
		wg.Done()
	}()

	go func() {
		for i := pivot;i <= n;i++{
			right *= i
		}
		wg.Done()
	}()

	wg.Wait()
	return left * right
}

// 发现型Go模式
func Find(n int) <-chan int{
	ch := make(chan int)
	go func() {
		prev := 1

		for i := 1; i <= n;i++{
			v := prev * i

			ch <- v
			prev = v
		}

		close(ch)
	}()
	return ch
}


//使用成熟的解决方案修复Go缺陷
type IFactorial interface {
	CalculateFactorial() int
}
var _ IFactorial = (*FactorialImpl)(nil)
type FactorialImpl struct {
	n int
}
func NewFactorial(n int) *FactorialImpl {
	return &FactorialImpl{
		n: n,
	}
}
func (this *FactorialImpl) GetN() int {
	return this.n
}
func (this *FactorialImpl) SetN(n int) {
	this.n = n
}
func (this *FactorialImpl) CalculateFactorial() int {
	if this.n == 0 {
		return 1
	}

	n := this.n
	this.n = this.n - 1

	return this.CalculateFactorial() * n
}


//高级Go程序员
func Senior(n int) int{
	res := 1

	for i := 1; i <= n;i++{
		res *= i
	}
	return res
}

func RobPike(n int) int{
	res := 1

	for i := 1;i <= n;i++{
		res *= i
	}

	return res
}

