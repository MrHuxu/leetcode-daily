\>\> [题目链接](https://leetcode.com/explore/featured/card/april-leetcoding-challenge-2021/596/week-4-april-22nd-april-28th/3722/)

题意: 给定一个数字 `n`, 判断是否是 3 的乘方

解答: 一个优化的思路是, 不要每次都是用 3 来除, 而是使用 3^1, 3^2, 3^4, 3^8... 这样的序列来除, 这里面的知识点是任何一个数都可以表示成若干个 2 的整次幂之和, 这样可以得到对数时间复杂度

***original_content***

<p>Given an integer <code>n</code>, return <em><code>true</code> if it is a power of three. Otherwise, return <code>false</code></em>.</p>

<p>An integer <code>n</code> is a power of three, if there exists an integer <code>x</code> such that <code>n == 3<sup>x</sup></code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> n = 27
<strong>Output:</strong> true
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> n = 0
<strong>Output:</strong> false
</pre><p><strong>Example 3:</strong></p>
<pre><strong>Input:</strong> n = 9
<strong>Output:</strong> true
</pre><p><strong>Example 4:</strong></p>
<pre><strong>Input:</strong> n = 45
<strong>Output:</strong> false
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= n &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> Could you solve it without loops/recursion?
