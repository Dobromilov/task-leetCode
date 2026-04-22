/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    var result []int
    if root == nil{
        return []int{}
    }
    leftPart := inorderTraversal(root.Left)
    currentPart := append(leftPart, root.Val)
    result = append(currentPart, inorderTraversal(root.Right)...)
    return result
}