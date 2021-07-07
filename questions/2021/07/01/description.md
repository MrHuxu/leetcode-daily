\>\> [题目链接](https://leetcode.com/explore/featured/card/july-leetcoding-challenge-2021/608/week-1-july-1st-july-7th/3799/)

题意: 给定数字 n, 生成 n 位的格雷码, 即 2^n 个 n 位的二进制数, 且相邻数只有一位不同

解答: 假设 n 位格雷码为 [i1,i2...i2^n], 那么对于 n+1, 我们把前面的序列中倒序添加到末尾, 即 [i1,i2...i2^n,i2^n...i2,i1], 然后前一半低位补 0, 后一半低位补 1, 就得到了 n+1 位的格雷码

***original_content***

<p>An <strong>n-bit gray code sequence</strong> is a sequence of <code>2<sup>n</sup></code> integers where:</p>

<ul>
	<li>Every integer is in the <strong>inclusive</strong> range <code>[0, 2<sup>n</sup> - 1]</code>,</li>
	<li>The first integer is <code>0</code>,</li>
	<li>An integer appears <strong>no more than once</strong> in the sequence,</li>
	<li>The binary representation of every pair of <strong>adjacent</strong> integers differs by <strong>exactly one bit</strong>, and</li>
	<li>The binary representation of the <strong>first</strong> and <strong>last</strong> integers differs by <strong>exactly one bit</strong>.</li>
</ul>

<p>Given an integer <code>n</code>, return <em>any valid <strong>n-bit gray code sequence</strong></em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 2
<strong>Output:</strong> [0,1,3,2]
<strong>Explanation:</strong>
The binary representation of [0,1,3,2] is [00,01,11,10].
- 0<u>0</u> and 0<u>1</u> differ by one bit
- <u>0</u>1 and <u>1</u>1 differ by one bit
- 1<u>1</u> and 1<u>0</u> differ by one bit
- <u>1</u>0 and <u>0</u>0 differ by one bit
[0,2,3,1] is also a valid gray code sequence, whose binary representation is [00,10,11,01].
- <u>0</u>0 and <u>1</u>0 differ by one bit
- 1<u>0</u> and 1<u>1</u> differ by one bit
- <u>1</u>1 and <u>0</u>1 differ by one bit
- 0<u>1</u> and 0<u>0</u> differ by one bit
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 1
<strong>Output:</strong> [0,1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 16</code></li>
</ul>

