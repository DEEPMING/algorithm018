/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if(l1 == nil){
        return l2
    }
    if(l2 == nil){
        return l1
    }

	var head ListNode
	var tail ListNode
	if l1.Val <= l2.Val {
		head.Next = l1
		tail.Next = l1
		l1 = l1.Next
	}else {
        head.Next = l2
		tail.Next = l2
		l2 = l2.Next
    }
	for true {
        if l1==nil && l2==nil {
			break
		}
        if(l1 == nil){
            tail.Next.Next =l2
			tail.Next = l2
			l2 = l2.Next
            continue
        }else if(l2 ==nil){
            tail.Next.Next =l1
			tail.Next = l1
			l1 = l1.Next
            continue
        }
        
		if l1.Val<=l2.Val {
			tail.Next.Next =l1
			tail.Next = l1
			l1 = l1.Next
		}else if l1.Val>l2.Val {
			tail.Next.Next =l2
			tail.Next = l2
			l2 = l2.Next
		}
		
	}

	return head.Next
}