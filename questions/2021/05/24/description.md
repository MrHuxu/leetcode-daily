\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/601/week-4-may-22nd-may-28th/3754/)

题意: 给定一个字符串 `s`, 将其转换为小写

解答: 解答可以直接用 `strings.ToLower`, 底层实现是声明一个 `strings.Builder` 存储转换结果, 然后逐个字符获得对应的小写字符然后使用 `WriteByte` 写入到结果里

***original_content***

<p>Given a string <code>s</code>, return <em>the string after replacing every uppercase letter with the same lowercase letter</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;Hello&quot;
<strong>Output:</strong> &quot;hello&quot;
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;here&quot;
<strong>Output:</strong> &quot;here&quot;
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;LOVELY&quot;
<strong>Output:</strong> &quot;lovely&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code> consists of printable ASCII characters.</li>
</ul>

