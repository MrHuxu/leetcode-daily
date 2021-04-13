\>\> [题目链接](https://leetcode.com/explore/featured/card/april-leetcoding-challenge-2021/594/week-2-april-8th-april-14th/3705/)

题意: 给定两个正整数 `n` 和 `k`, 返回一个长度为 n 的数组 arr 包含 1~n 的数字, 并且满足 `|arr[0]-arr[1]|`, `|arr[1]-arr[2]|` ... `|arr[n-2]-arr[n-1]|` 满足有 k 个不同的值.

解答: 前 k + 1 个数字使用这样的方式来构建

[1, 1 + k, 2, k, 3, k - 1 ...]

后面直接把 k+2~n 个数字填到数组末尾即可.