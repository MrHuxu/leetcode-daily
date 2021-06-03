\>\> [题目链接](https://leetcode.com/explore/featured/card/june-leetcoding-challenge-2021/603/week-1-june-1st-june-7th/3766/)

题意: 给定一个蛋糕的长 `h` 和宽 `w`, 然后给定一个数组 `horizontalCuts` 表示水平切蛋糕的位置, 数组 `verticalCuts` 表示垂直切蛋糕的位置, 求切完之后最大的蛋糕的面积

解答: 遍历找出切完之后最大的长和宽返回乘积即可

***original_content***

<p>Given a rectangular cake with height <code>h</code> and width <code>w</code>, and two arrays of integers <code>horizontalCuts</code> and <code>verticalCuts</code> where <code>horizontalCuts[i]</code> is the distance from the top of the rectangular cake to the <code>ith</code> horizontal cut&nbsp;and similarly, <code>verticalCuts[j]</code> is the distance from the&nbsp;left of the rectangular cake to the <code>jth</code>&nbsp;vertical cut.</p>

<p><em>Return the maximum area of a piece of cake after you cut at each horizontal and vertical position provided in the arrays <code>horizontalCuts</code> and <code>verticalCuts</code>.&nbsp;</em>Since the answer can be a huge number, return this modulo 10^9 + 7.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2020/05/14/leetcode_max_area_2.png" style="width: 300px; height: 320px;" /></p>

<pre>
<strong>Input:</strong> h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
<strong>Output:</strong> 4 
<strong>Explanation:</strong> The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green piece of cake has the maximum area.
</pre>

<p><strong>Example 2:</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2020/05/14/leetcode_max_area_3.png" style="width: 300px; height: 320px;" /></strong></p>

<pre>
<strong>Input:</strong> h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
<strong>Output:</strong> 6
<strong>Explanation:</strong> The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green and yellow pieces of cake have the maximum area.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
<strong>Output:</strong> 9
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= h,&nbsp;w &lt;= 10^9</code></li>
	<li><code>1 &lt;=&nbsp;horizontalCuts.length &lt;&nbsp;min(h, 10^5)</code></li>
	<li><code>1 &lt;=&nbsp;verticalCuts.length &lt; min(w, 10^5)</code></li>
	<li><code>1 &lt;=&nbsp;horizontalCuts[i] &lt; h</code></li>
	<li><code>1 &lt;=&nbsp;verticalCuts[i] &lt; w</code></li>
	<li>It is guaranteed that all elements in&nbsp;<code>horizontalCuts</code>&nbsp;are distinct.</li>
	<li>It is guaranteed that all elements in <code>verticalCuts</code>&nbsp;are distinct.</li>
</ul>

