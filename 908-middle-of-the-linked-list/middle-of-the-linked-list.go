/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    l := 1
    curl := head

    for curl.Next != nil {
        l++
        curl = curl.Next
    }
    
    if l == 1{
        return head
    }
    
    for i:=0; i<l/2; i++{
        head = head.Next
    }
    return head
}