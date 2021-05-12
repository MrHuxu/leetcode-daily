\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/599/week-2-may-8th-may-14th/3737/)

题意:	给定一个数组 `target`, 每一次可以对一个数组做如下操作:

1. 求数组的和
2. 把数组的任意一位置为数组和的值

判断初始数字全都为 1 的话, target 是否可以通过上面的操作得到

解答: 每一步找出数组中的最大值, 这个值肯定是上一次操作之后的和, 然后就可以求出上一次变化之后的数组, 然后递归求解即可

***original_content***

<p>Given an array of integers&nbsp;<code>target</code>. From a starting array, <code>A</code>&nbsp;consisting of all 1&#39;s, you may perform the following procedure :</p>

<ul>
	<li>let <code>x</code> be the sum of all elements currently in your array.</li>
	<li>choose index <code>i</code>, such that&nbsp;<code>0 &lt;= i &lt; target.size</code> and set the value of <code>A</code> at index <code>i</code> to <code>x</code>.</li>
	<li>You may repeat this procedure&nbsp;as many times as needed.</li>
</ul>

<p>Return True if it is possible to construct the <code>target</code> array from <code>A</code> otherwise&nbsp;return False.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> target = [9,3,5]
<strong>Output:</strong> true
<strong>Explanation:</strong> Start with [1, 1, 1] 
[1, 1, 1], sum = 3 choose index 1
[1, 3, 1], sum = 5 choose index 2
[1, 3, 5], sum = 9 choose index 0
[9, 3, 5] Done
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> target = [1,1,1,2]
<strong>Output:</strong> false
<strong>Explanation:</strong> Impossible to create target array from [1,1,1,1].
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> target = [8,5]
<strong>Output:</strong> true
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>N == target.length</code></li>
	<li><code>1 &lt;= target.length&nbsp;&lt;= 5 * 10^4</code></li>
	<li><code>1 &lt;= target[i] &lt;= 10^9</code></li>
</ul>

