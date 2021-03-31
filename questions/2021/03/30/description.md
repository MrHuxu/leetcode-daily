\>\> [题目链接](https://leetcode.com/explore/featured/card/march-leetcoding-challenge-2021/592/week-5-march-29th-march-31st/3690/)

题意: 给一个数组 `envelopes`, 每一项是一个两元组表示一个信封的长和宽, 当一个信封的长宽都大于另一个信封时, 这两个信封可以套起来, 求可以套起来最大信封数

解答: 即求长宽都严格递增的最大子数组大小, 可以通过长进行排序, 然后使用一个长度为 `len(envelopes)` 的数组 `dp` 来记录每一项可以套起来信封数, 每遍历到一项, 就往前回溯, 如果存在 `envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1]`, 则信封 j 可以被信封 i 套住, 使用 `dp[i] = max(dp[i], dp[j] + 1)` 来更新值, 同时更新结果为当前已有的最大值, 遍历结束返回即可