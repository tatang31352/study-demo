package twoSum

import "testing"

func TestIsSet(t *testing.T){
	cas := []struct{
		nums []int
		value int
		key int
		res int
	}{
		{[]int{1,2,3,4,5,6,7},6,2,5},
		{[]int{1,2,3,4,5,6,7},8,2,0},
		{[]int{1,2,3,4,5,6,7},3,2,0},
	}

	for _,c := range cas{
		if c.res != IsSet(c.nums,c.value,c.key){
			t.Fatalf("isSet function failed, nums:%d,value:%d,key:%d,res:%d",c.nums,c.value,c.key,c.key)
		}
	}
}

func TestTwoSum(t *testing.T){
	cas := []struct{
		nums []int
		target int
		res []int
	}{
		{[]int{1,2,3,4,5},9,[]int{3,4}},
		{[]int{1,2,3,4,5,7,8},15,[]int{5,6}},
		{[]int{1,2,3,4,5,10,6,8},16,[]int{5,6}},
		{[]int{1,2,3,4,5},3,[]int{0,1}},
	}

	var result []int
	for _,c := range cas{

		result =  TwoSum(c.nums,c.target)
		if c.res[0] != result[0] ||  c.res[1] != result[1]{
			t.Fatalf("isSet function failed, nums:%d,targer:%d,res:%d",c.nums,c.target,c.res)
		}
	}

}