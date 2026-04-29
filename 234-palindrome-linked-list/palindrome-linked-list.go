/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseInts(input []int) []int {
    newArr := make([]int, len(input))
    for i, j := 0, len(input)-1; i < len(input); i, j = i+1, j-1 {
        newArr[i] = input[j]
    }
    return newArr
}

func isPalindrome(head *ListNode) bool {
    curl := head
    array := make([]int, 0, 1)
    if head == nil{
        return false
    }

    for curl.Next != nil {
        array = append(array,curl.Val)
        curl = curl.Next
    }

    array = append(array,curl.Val)
    temp := reverseInts(array)
    
    for i:=0; i<len(array); i++{
        if array[i]!=temp[i]{
            return false
        }
    }
    return true
}