\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/599/week-2-may-8th-may-14th/3741/)

题意: 给定一个字符串 `s`, 将其中包含的数字变成两个合法数字的二元组的形式

解答: 首先遍历 s 得到逗号插入的位置, 然后对字符串的左右两边进行处理, 也就是根据首尾的 0 来获得合法的数字列表, 然后对左右的合法数字列表进行组合, 就可以得到最终结果

***original_content***

<p>We had some 2-dimensional coordinates, like <code>&quot;(1, 3)&quot;</code> or <code>&quot;(2, 0.5)&quot;</code>.&nbsp; Then, we removed&nbsp;all commas, decimal points, and spaces, and ended up with the string&nbsp;<code>s</code>.&nbsp; Return a list of strings representing&nbsp;all possibilities for what our original coordinates could have been.</p>

<p>Our original representation never had extraneous zeroes, so we never started with numbers like &quot;00&quot;, &quot;0.0&quot;, &quot;0.00&quot;, &quot;1.0&quot;, &quot;001&quot;, &quot;00.01&quot;, or any other number that can be represented with&nbsp;less digits.&nbsp; Also, a decimal point within a number never occurs without at least one digit occuring before it, so we never started with numbers like &quot;.1&quot;.</p>

<p>The final answer list can be returned in any order.&nbsp; Also note that all coordinates in the final answer&nbsp;have exactly one space between them (occurring after the comma.)</p>

<pre>
<strong>Example 1:</strong>
<strong>Input:</strong> s = &quot;(123)&quot;
<strong>Output:</strong> [&quot;(1, 23)&quot;, &quot;(12, 3)&quot;, &quot;(1.2, 3)&quot;, &quot;(1, 2.3)&quot;]
</pre>

<pre>
<strong>Example 2:</strong>
<strong>Input:</strong> s = &quot;(00011)&quot;
<strong>Output:</strong> &nbsp;[&quot;(0.001, 1)&quot;, &quot;(0, 0.011)&quot;]
<strong>Explanation:</strong> 
0.0, 00, 0001 or 00.01 are not allowed.
</pre>

<pre>
<strong>Example 3:</strong>
<strong>Input:</strong> s = &quot;(0123)&quot;
<strong>Output:</strong> [&quot;(0, 123)&quot;, &quot;(0, 12.3)&quot;, &quot;(0, 1.23)&quot;, &quot;(0.1, 23)&quot;, &quot;(0.1, 2.3)&quot;, &quot;(0.12, 3)&quot;]
</pre>

<pre>
<strong>Example 4:</strong>
<strong>Input:</strong> s = &quot;(100)&quot;
<strong>Output:</strong> [(10, 0)]
<strong>Explanation:</strong> 
1.0 is not allowed.
</pre>

<p>&nbsp;</p>

<p><strong>Note: </strong></p>

<ul>
	<li><code>4 &lt;= s.length &lt;= 12</code>.</li>
	<li><code>s[0]</code> = &quot;(&quot;, <code>s[s.length - 1]</code> = &quot;)&quot;, and the other elements in <code>s</code> are digits.</li>
</ul>

<p>&nbsp;</p>

