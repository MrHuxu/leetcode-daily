\>\> [题目链接](https://leetcode.com/explore/featured/card/june-leetcoding-challenge-2021/603/week-1-june-1st-june-7th/3769/)

题意: 给定一个整数数组 `nums`, 找出里面最长连续子数组, 要求时间复杂度为 O(n)

解答: 维护两个 map, `m1 map[int]int` 使用 m1[i][j] 表示存在从 i 到 j 的连续递增子数组, `m2 map[int]int` 使用 m2[i][j] 表示存在从 i 到 j 的连续递减子数组, 遍历 nums 使用这两个 map 来更新最长连续子数组的首尾位置, 最后可以返回遍历过程中的最大值

***original_content***

<p>Given an unsorted array of integers <code>nums</code>, return <em>the length of the longest consecutive elements sequence.</em></p>

<p>You must write an algorithm that runs in&nbsp;<code>O(n)</code>&nbsp;time.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [100,4,200,1,3,2]
<strong>Output:</strong> 4
<strong>Explanation:</strong> The longest consecutive elements sequence is <code>[1, 2, 3, 4]</code>. Therefore its length is 4.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,3,7,2,5,8,4,6,0,1]
<strong>Output:</strong> 9
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

