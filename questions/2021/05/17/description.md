\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/600/week-3-may-15th-may-21st/3746/)

题意: 给定一个单词列表 `words`, 如果一个单词 word1 在任意位置加一个字母可以获得单词 word2, 则称 word1 是 word2 的前辈, 判断 words 里最长的前辈链

解答: 先按照长度对 words 做分组, 然后从 1 到最大长度遍历, 每遍历一个分组, 就判断长度减 1 的分组里有没有前辈, 并且更新前辈链的长度, 最后返回最长长度即可

***original_content***

<p>Given a list of words, each word consists of English lowercase letters.</p>

<p>Let&#39;s say <code>word1</code> is a predecessor of <code>word2</code> if and only if we can add exactly one letter anywhere in <code>word1</code> to make it equal to <code>word2</code>. For example, <code>&quot;abc&quot;</code> is a predecessor of <code>&quot;abac&quot;</code>.</p>

<p>A <em>word chain </em>is a sequence of words <code>[word_1, word_2, ..., word_k]</code> with <code>k &gt;= 1</code>, where <code>word_1</code> is a predecessor of <code>word_2</code>, <code>word_2</code> is a predecessor of <code>word_3</code>, and so on.</p>

<p>Return the longest possible length of a word chain with words chosen from the given list of <code>words</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;a&quot;,&quot;b&quot;,&quot;ba&quot;,&quot;bca&quot;,&quot;bda&quot;,&quot;bdca&quot;]
<strong>Output:</strong> 4
<strong>Explanation</strong>: One of the longest word chain is &quot;a&quot;,&quot;ba&quot;,&quot;bda&quot;,&quot;bdca&quot;.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> words = [&quot;xbc&quot;,&quot;pcxbcf&quot;,&quot;xb&quot;,&quot;cxbc&quot;,&quot;pcxbc&quot;]
<strong>Output:</strong> 5
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 1000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 16</code></li>
	<li><code>words[i]</code> only consists of English lowercase letters.</li>
</ul>

