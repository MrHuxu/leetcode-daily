\>\> [题目链接](https://leetcode.com/explore/featured/card/april-leetcoding-challenge-2021/596/week-4-april-22nd-april-28th/3718/)

题意: 给定一个二进制字符串 `s`, 找出里面包含 1 和 0 个数相等的子串, 并且子串里 1 和 0 都是连续的

解答: 三种解法, 效率逐个提高

1. dp: 声明二维数组 `dp [][]bool` 使用 dp[i][l] 表示起始位置为 i 长度为 l 的子串是不是符合要求的子串, 那么有这样的状态转移方程 `dp[i][l] = dp[i+1][l-2] && (s[i] != s[i+l-1] && s[i] == s[i+1] && s[i+l-1] == s[i+l-2])`
2. 遍历解法1: 因为 1 和 0 都是连续的, 所以只需要在遍历到时候, 找到 1 和 0 相邻的位置, 然后向前后继续遍历, 就可以得到这一段子数组包含的合法子串数量
3. 遍历解法2: 保存上一个遍历到字符 `pre` 和字符数量 `preCnt`, 以及当前遍历到字符 `curr` 和字符数量 `currCnt`, 这样的话只需要当遍历的字符不等于 curr, 把 preCnt 和 currCnt 的较小值累加到结果, 使用 curr 和 currCnt 更新 pre 和 preCnt, 并且将 curr 更新为当前字符即可, 这样的复杂度是 O(n), 是最快的解法

***original_content***

<p>Give a string <code>s</code>, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively. 
</p>
<p>Substrings that occur multiple times are counted the number of times they occur.</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> "00110011"
<b>Output:</b> 6
<b>Explanation:</b> There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
<br>Notice that some of these substrings repeat and are counted the number of times they occur.
<br>Also, "00110011" is not a valid substring because <b>all</b> the 0's (and 1's) are not grouped together.
</pre>
</p>

<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> "10101"
<b>Output:</b> 4
<b>Explanation:</b> There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
</pre>
</p>

<p><b>Note:</b>
<li><code>s.length</code> will be between 1 and 50,000.</li>
<li><code>s</code> will only consist of "0" or "1" characters.</li>
</p>
