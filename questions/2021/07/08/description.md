\>\> [题目链接](https://leetcode.com/explore/featured/card/july-leetcoding-challenge-2021/609/week-2-july-8th-july-14th/3807/)

题意: 给定两个数组 `nums1` 和 `nums2`, 求两个数组的相同子数组的最大长度

解答: 声明二维数组 `dp [][]int` 使用 dp[i][j] 表示包含 nums1 的第 i 个数字和 nums2 的第 j 个数字的相同子数组的最大长度, 那么有如下的状态转移方程:

1. i 或 j 为 0, dp[i][j] 为 0
2. nums1[i-1] != nums2[j-1], dp[i][j] 为 0
3. nums1[i-1] == nums2[j-1], dp[i][j] = dp[i-1][j-1]+1

同时更新结果为 dp[i][j] 计算过程中的最大值, 最后返回即可

***original_content***

<p>Given two integer arrays <code>nums1</code> and <code>nums2</code>, return <em>the maximum length of a subarray that appears in <strong>both</strong> arrays</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
<strong>Output:</strong> 3
<strong>Explanation:</strong> The repeated subarray with maximum length is [3,2,1].
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
<strong>Output:</strong> 5
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length, nums2.length &lt;= 1000</code></li>
	<li><code>0 &lt;= nums1[i], nums2[i] &lt;= 100</code></li>
</ul>

