func moveZeroes(nums []int)  {
    index := 0
    for _,i := range nums {
        if i != 0 {
            //fmt.Println(i)
            nums[index] = i
            index ++
        }
    }
    for ;index < len(nums); index ++ {
        nums[index] = 0
    }
}