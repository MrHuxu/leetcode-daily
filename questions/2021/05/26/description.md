\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/601/week-4-may-22nd-may-28th/3756/)

题意: 给定一个十进制数字的字符串 `n`, 这个数字最少可以由几个只由 `0` 和 `1` 组成的数字相加得到

解答: 因为每个数字最多给一位累加 1, 所以需要的数量和 n 各位上最大的数字一致, 直接遍历即可, 用堆做的话因为增加了数字交换的性能损失, 耗时反而会增加

***original_content***

<p>A decimal number is called <strong>deci-binary</strong> if each of its digits is either <code>0</code> or <code>1</code> without any leading zeros. For example, <code>101</code> and <code>1100</code> are <strong>deci-binary</strong>, while <code>112</code> and <code>3001</code> are not.</p>

<p>Given a string <code>n</code> that represents a positive decimal integer, return <em>the <strong>minimum</strong> number of positive <strong>deci-binary</strong> numbers needed so that they sum up to </em><code>n</code><em>.</em></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = &quot;32&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> 10 + 11 + 11 = 32
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = &quot;82734&quot;
<strong>Output:</strong> 8
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> n = &quot;27346209830709182346&quot;
<strong>Output:</strong> 9
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n.length &lt;= 10<sup>5</sup></code></li>
	<li><code>n</code> consists of only digits.</li>
	<li><code>n</code> does not contain any leading zeros and represents a positive integer.</li>
</ul>
