\>\> [题目链接](https://leetcode.com/explore/featured/card/april-leetcoding-challenge-2021/594/week-2-april-8th-april-14th/3703/)

题意: 给定一个二维数组 `matrix`, 从每一个点只能向上下左右四个方向移动, 并且不能移动出边界, 找出最长递增路径的长度.

解答: dfs+dp, 使用一个二维数组 `dp [][]int` 使用 dp[i][j] 表示以 matrix[i][j] 为起点的最长递增路径长度, 避免对一个位置反复求解, 使用 dfs 对四个方向遍历求解即可