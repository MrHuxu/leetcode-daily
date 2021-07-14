\>\> [题目链接](https://leetcode.com/explore/featured/card/july-leetcoding-challenge-2021/609/week-2-july-8th-july-14th/3813/)

题意: 给定一个顺序字符串 `order` 和一个源字符串 `str`, 将 str 里的字符按照 order 里出现的顺序进行排序

解答: 先将 order 里字符和索引做一个映射, 然后实现 sort 函数即可

***original_content***

<p><code>order</code> and <code>str</code> are strings composed of lowercase letters. In <code>order</code>, no letter occurs more than once.</p>

<p><code>order</code> was sorted in some custom order previously. We want to permute the characters of <code>str</code> so that they match the order that <code>order</code> was sorted. More specifically, if <code>x</code> occurs before <code>y</code> in <code>order</code>, then <code>x</code> should occur before <code>y</code> in the returned string.</p>

<p>Return any permutation of <code>str</code> (as a string) that satisfies this property.</p>

<pre>
<strong>Example:</strong>
<strong>Input:</strong> 
order = &quot;cba&quot;
str = &quot;abcd&quot;
<strong>Output:</strong> &quot;cbad&quot;
<strong>Explanation:</strong> 
&quot;a&quot;, &quot;b&quot;, &quot;c&quot; appear in order, so the order of &quot;a&quot;, &quot;b&quot;, &quot;c&quot; should be &quot;c&quot;, &quot;b&quot;, and &quot;a&quot;. 
Since &quot;d&quot; does not appear in order, it can be at any position in the returned string. &quot;dcba&quot;, &quot;cdba&quot;, &quot;cbda&quot; are also valid outputs.
</pre>

<p>&nbsp;</p>

<p><strong>Note:</strong></p>

<ul>
	<li><code>order</code> has length at most <code>26</code>, and no character is repeated in <code>order</code>.</li>
	<li><code>str</code> has length at most <code>200</code>.</li>
	<li><code>order</code> and <code>str</code> consist of lowercase letters only.</li>
</ul>

