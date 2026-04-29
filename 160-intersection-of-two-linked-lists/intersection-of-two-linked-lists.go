/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]bool)
    
    currA := headA
    for currA != nil {
        m[currA] = true
        currA = currA.Next
    }
    currB := headB
    for currB != nil {
        if m[currB] {
            return currB
        }
        currB = currB.Next
    }
    
    return nil
}
