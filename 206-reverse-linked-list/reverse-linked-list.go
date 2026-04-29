/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    array := make([]int, 0, 1)
    curl := head

    if head == nil{
        return head
    }

    for curl.Next != nil {
        array = append(array,curl.Val)
        curl = curl.Next
    }
    array = append(array,curl.Val)
    slices.Reverse(array)
    
    res := head
    for i := 0; i < len(array); i++ {
        res.Val = array[i]
        res = res.Next
    }
    return head
}