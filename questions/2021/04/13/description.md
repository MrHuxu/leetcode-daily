\>\> [题目链接](https://leetcode.com/explore/featured/card/april-leetcoding-challenge-2021/594/week-2-april-8th-april-14th/3706/)

题意: 打平一组嵌套数组并实现带遍历函数 `Next` 的迭代器.

解答: 使用递归将输入数据存入一个打平的数组里, 然后记录遍历位置 `idx`, 调用 Next 时累加 idx 即可.