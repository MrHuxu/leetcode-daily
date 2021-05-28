\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/601/week-4-may-22nd-may-28th/3757/)

题意: 给定一组单词 `words`, 判断不存在相同字母的两个单词长度的乘积的最大值

解答: 使用位图存储一个单词里一个字母是否出现, 然后就可以快速判断两个单词是否存在相同字母, 最后遍历找出最大乘积即可

***original_content***

<p>Given a string array <code>words</code>, return <em>the maximum value of</em> <code>length(word[i]) * length(word[j])</code> <em>where the two words do not share common letters</em>. If no such two words exist, return <code>0</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;abcw&quot;,&quot;baz&quot;,&quot;foo&quot;,&quot;bar&quot;,&quot;xtfn&quot;,&quot;abcdef&quot;]
<strong>Output:</strong> 16
<strong>Explanation:</strong> The two words can be &quot;abcw&quot;, &quot;xtfn&quot;.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;a&quot;,&quot;ab&quot;,&quot;abc&quot;,&quot;d&quot;,&quot;cd&quot;,&quot;bcd&quot;,&quot;abcd&quot;]
<strong>Output:</strong> 4
<strong>Explanation:</strong> The two words can be &quot;ab&quot;, &quot;cd&quot;.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;a&quot;,&quot;aa&quot;,&quot;aaa&quot;,&quot;aaaa&quot;]
<strong>Output:</strong> 0
<strong>Explanation:</strong> No such pair of words.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= words.length &lt;= 1000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 1000</code></li>
	<li><code>words[i]</code> consists only of lowercase English letters.</li>
</ul>

