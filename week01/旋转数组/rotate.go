func rotate(nums []int, k int)  {
	k = k % len(nums)
	if k == 0 {
		return
	}
	rev(nums)
	rev(nums[:k])
	rev(nums[k:])

}

func rev(nums []int) {
	first :=0
	last := len(nums)-1
	for true {
		nums[first], nums[last] = nums[last], nums[first]
		first++
		last--
		if first == last || first>last {
			break
		}
	}
}