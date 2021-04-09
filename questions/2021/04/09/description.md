\>\> [题目链接](https://leetcode.com/explore/featured/card/april-leetcoding-challenge-2021/594/week-2-april-8th-april-14th/3702/)

题意: 给一个字符串数组 `words`, 和一个字符串 `order`, 判断 words 里的字符串是否是按照 order 里的字母先后顺序进行排序.

解答: 首先将 order 里的内容存到 `letterOrder` 里方便后面获取一个字母的顺序, 然后从第一个字符开始遍历 words 里相邻的字符串:

1. 如果存在不符合 order 顺序的情况, 直接返回 false
2. 如果存在严格遵守 order 顺序(即相邻的字符串里的字母严格按 order 顺序并且不相等)的情况, 当前判断为 true, 继续下一次
3. 如果没走到上面两种情况, 那么一个字符串为另一个的前缀, 这时短的那个应该在前面, 在后面的话返回 false

最后返回 true 即可.