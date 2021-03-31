\>\> [题目链接](https://leetcode.com/explore/featured/card/march-leetcoding-challenge-2021/592/week-5-march-29th-march-31st/3689/)

题意: 给定一个二叉树根节点 `root` 以及一个遍历序列 `voyage`, 判断是否可以通过翻转二叉树的一些节点的左右子节点使其前序遍历结果为 voyage, 如果可以这样的操作话, 返回所有需要翻转的节点值, 否则返回 `[-1]`

解答: 二叉树的前序遍历序列, 根节点一定在序列的开始位置, 根据这个规律递归遍历判断即可, 注意区分左右子节点不存在的情况