type MyCircularDeque struct {
    nums []int
    length int
    size int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        nums : make([]int, 0, 10),
        size : k,
        length : 0,
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.length < this.size {
        this.nums = append([]int{value}, this.nums...)
        this.length ++
        return true
    }
    return false
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.length < this.size {
        this.nums = append(this.nums, value)
        this.length ++
        return true
    }
    return false
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.length > 0 {
        this.nums = this.nums[1:]
        this.length --
        return true
    }
    return false
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.length > 0 {
        this.nums = this.nums[:len(this.nums)-1]
        this.length --
        return true
    }
    return false
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.length > 0 {
        return this.nums[0]
        
    }
    return -1
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.length > 0 {
        return this.nums[len(this.nums)-1]
    }
    return -1
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    if this.length == 0 {
        //fmt.Println(this.length, this.size, "Is Empty")
        return true
    }
    return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    if this.length >= this.size {
        //fmt.Println(this.length, this.size, "Is Full")
        return true
    }
    return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */