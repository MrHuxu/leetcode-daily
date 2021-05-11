\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/599/week-2-may-8th-may-14th/3736/)

题意: 字符串 `left` 和 `right` 表示左右边界, 找到左右边界里的超级回文数, 即自身和平方根都是回文数

解答: 直接遍历的话复杂度非常高, 可以使用一个函数来生成回文数, 然后判断数字是否在 `math.Ceil(math.Sqrt(left))` 和 `math.Floor(math.Sqrt(right))` 内, 并且数字的平方也是回文数, 如果都符合的话对结果累加

***original_content***

<p>Let&#39;s say a positive integer is a <strong>super-palindrome</strong> if it is a palindrome, and it is also the square of a palindrome.</p>

<p>Given two positive integers <code>left</code> and <code>right</code> represented as strings, return <em>the number of <strong>super-palindromes</strong> integers in the inclusive range</em> <code>[left, right]</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> left = &quot;4&quot;, right = &quot;1000&quot;
<strong>Output:</strong> 4
<strong>Explanation</strong>: 4, 9, 121, and 484 are superpalindromes.
Note that 676 is not a superpalindrome: 26 * 26 = 676, but 26 is not a palindrome.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> left = &quot;1&quot;, right = &quot;2&quot;
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= left.length, right.length &lt;= 18</code></li>
	<li><code>left</code> and <code>right</code> consist of only digits.</li>
	<li><code>left</code> and <code>right</code> cannot have leading zeros.</li>
	<li><code>left</code> and <code>right</code> represent integers in the range <code>[1, 10<sup>18</sup>]</code>.</li>
	<li><code>left</code> is less than or equal to <code>right</code>.</li>
</ul>

