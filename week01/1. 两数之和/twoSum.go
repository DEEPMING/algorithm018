func twoSum(nums []int, target int) []int {
    //双重循环
    // slice := make([]int,2)
    // for i := 0; i<len(nums)-1; i++ {
    //     for j := i+1; j<len(nums); j++ {
    //         if (nums[i]+nums[j]) == target {
    //             slice[0] =i
    //             slice[1] =j
    //             return slice
    //         }
    //     }
    // }
    // return slice

    m := make(map[int]int)
    for k, v := range nums {
        m[v] = k
    }
    for k, v := range nums {
        if m[target-v] !=0 && m[target-v]!=k {
            return []int{k, m[target-v]}
        }
    }

    return []int{}
}