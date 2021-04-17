\>\> [题目连接](https://leetcode.com/explore/featured/card/march-leetcoding-challenge-2021/591/week-4-march-22nd-march-28th/3682/)

题意: 给定一个数组 `arr`, 找出里面所有和为 `target` 的三项组合数, 不同下标相同的值按不同组处理

解答: 先排序, 然后从前往后遍历, 同时对剩下的两项分别从头 `left` 和尾 `right` 向中间逼近, 注意需要处理值相同的情况

1. 如果此时三项和为 target, 则直接统计 `arr[left:right+1]` 的所有两元组
2. 如果此时三项和不为 target, 则没有必要继续逼近, 直接结束即可