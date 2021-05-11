\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/598/week-1-may-1st-may-7th/3728/)

题意: 设计一个特殊的字典, 给定一个字符串数组 `words` 表示字典里的单词, 实现一个函数 `f(prefix, suffix)` 找出字典里满足前缀和后缀的最大数组下标

解答: 前缀和后缀建立两棵 trie 树, 分别遍历得到结果, 选最大的即可

***original_content***

<p>Design a special dictionary with some words that searchs the words in it by a prefix and a suffix.</p>

<p>Implement the <code>WordFilter</code> class:</p>

<ul>
	<li><code>WordFilter(string[] words)</code> Initializes the object with the <code>words</code> in the dictionary.</li>
	<li><code>f(string prefix, string suffix)</code> Returns <em>the index of the word in the dictionary,</em> which has the prefix <code>prefix</code> and the suffix <code>suffix</code>. If there is more than one valid index, return <strong>the largest</strong> of them. If there is no such word in the dictionary, return <code>-1</code>.</li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input</strong>
[&quot;WordFilter&quot;, &quot;f&quot;]
[[[&quot;apple&quot;]], [&quot;a&quot;, &quot;e&quot;]]
<strong>Output</strong>
[null, 0]

<strong>Explanation</strong>
WordFilter wordFilter = new WordFilter([&quot;apple&quot;]);
wordFilter.f(&quot;a&quot;, &quot;e&quot;); // return 0, because the word at index 0 has prefix = &quot;a&quot; and suffix = &#39;e&quot;.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 15000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 10</code></li>
	<li><code>1 &lt;= prefix.length, suffix.length &lt;= 10</code></li>
	<li><code>words[i]</code>, <code>prefix</code> and <code>suffix</code> consist of lower-case English letters only.</li>
	<li>At most <code>15000</code> calls will be made to the function <code>f</code>.</li>
</ul>

