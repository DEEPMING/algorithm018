func trap(height []int) int {
    if len(height) < 2 {
        return 0
    }
    var left, right, leftMax, rightMax, result int
    left = 0
    right = len(height)-1
    leftMax = height[0]
    rightMax = height[right]

    for left < right {
        //fmt.Println(left, leftMax, right, rightMax, result)
        if height[left] < height[right] {
            if height[left] >=leftMax {
                leftMax = height[left]
            }else {
                result += leftMax - height[left]
            }
            left ++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                result += rightMax - height[right]
            }
            right --
        }
    }

    return result

}