\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/598/week-1-may-1st-may-7th/3731/)

题意: 给定一个整数数组 `nums`, 判断是否可以通过修改数组里的一个数字使其变成非递减序列

解答: 使用一个变量 `modified` 表示数组是否修改过, 那么遍历到 nums[i] 有如下几种情况:

1. i==0 或者 nums[i]>=nums[i-1], 直接往后遍历, 不破坏递增
2. nums[i]\<nums[i-1], 如果 modified=true 即数组已经被修改, 返回 false
3. nums[i]\<nums[i-1], 如果 i==1 或者 nums[i-2]<=nums[i], 那么直接将 nums[i-1] 向下改为 nums\[i\](**不将 nums[i]向上修改为 nums[i-1] 是为了不破坏 nums[i] 后面的递增**), 置 modified 为 true
4. nums[i]\<nums[i-1], 且不满足 3, 那么只能将 nums[i] 向上修改为 nums[i-1] 了, 置 modified 为 true

遍历结束跳出循环返回 true

***original_content***

<p>Given an array <code>nums</code> with <code>n</code> integers, your task is to check if it could become non-decreasing by modifying <strong>at most one element</strong>.</p>

<p>We define an array is non-decreasing if <code>nums[i] &lt;= nums[i + 1]</code> holds for every <code>i</code> (<strong>0-based</strong>) such that (<code>0 &lt;= i &lt;= n - 2</code>).</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [4,2,3]
<strong>Output:</strong> true
<strong>Explanation:</strong> You could modify the first <code>4</code> to <code>1</code> to get a non-decreasing array.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [4,2,1]
<strong>Output:</strong> false
<strong>Explanation:</strong> You can&#39;t get a non-decreasing array by modify at most one element.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>n == nums.length</code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>4</sup></code></li>
	<li><code>-10<sup>5</sup> &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>

