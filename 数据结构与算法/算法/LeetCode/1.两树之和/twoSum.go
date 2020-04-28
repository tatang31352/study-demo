package twoSum



func TwoSum(nums []int, target int) []int {
	for k,v := range nums{
		if IsSet(nums,target - v,k) != 0{
			return []int{k,IsSet(nums,target - v,k)}
		}
	}

	return []int{0,0}
}

func IsSet(nums []int,value int,key int) int{

	for k,v := range nums{
		if k != key{
			if v == value{
				return k
			}
		}
	}
	return 0
}
