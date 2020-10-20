func plusOne(digits []int) []int {
	var (
		head =1
	)
	for i := len(digits) - 1; i >= 0; i --{
		head, digits[i] = (digits[i]+head)/10, (digits[i]+head)%10
		if head == 0 {
			return digits
		}
	}
	slice := make([]int,1)
	slice[0] = 1
	slice = append(slice,digits...)
	return slice
}