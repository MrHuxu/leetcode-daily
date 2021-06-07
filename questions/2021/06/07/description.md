\>\> [题目链接](https://leetcode.com/explore/featured/card/june-leetcoding-challenge-2021/603/week-1-june-1st-june-7th/3770/)

题意: 给定一组整数数组 `cost` 表示跨每级台阶需要的代价, 每次可以向上跨一级或两级, 从位置 0 或者 1 开始, 求跨到顶需要的最少代价

解答: 简单 DP

***original_content***

<p>You are given an integer array <code>cost</code> where <code>cost[i]</code> is the cost of <code>i<sup>th</sup></code> step on a staircase. Once you pay the cost, you can either climb one or two steps.</p>

<p>You can either start from the step with index <code>0</code>, or the step with index <code>1</code>.</p>

<p>Return <em>the minimum cost to reach the top of the floor</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> cost = [10,15,20]
<strong>Output:</strong> 15
<strong>Explanation:</strong> Cheapest is: start on cost[1], pay that cost, and go to the top.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> cost = [1,100,1,1,1,100,1,1,100,1]
<strong>Output:</strong> 6
<strong>Explanation:</strong> Cheapest is: start on cost[0], and only step on 1s, skipping cost[3].
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= cost.length &lt;= 1000</code></li>
	<li><code>0 &lt;= cost[i] &lt;= 999</code></li>
</ul>

