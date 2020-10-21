func removeDuplicates(nums []int) int {
	
	index := 1
	for _,v := range nums{
		if k == 0 {
			continue
		}
		if v != nums[k-1] {
			nums[index], index = v, index+1
		}
	}
	
	return index
}