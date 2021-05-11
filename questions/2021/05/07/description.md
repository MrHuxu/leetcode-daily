\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/598/week-1-may-1st-may-7th/3734/)

题意: 给定两个单词 `word1` 和 `word2`, 每次操作可以删除任意一个单词里的一个字符, 判断最少操作几次可以使两个单词一致

解答: 声明二维数组 `dp [][]int` 使用 dp[i][j] 表示 word1的前 i 个字符和 word2 的前 j 个字符最少可以通过几次删除操作达到一致, 状态转移方程很好推, 遍历然后返回 dp[len(word1)][len(word2)] 即可

***original_content***

<p>Given two strings <code>word1</code> and <code>word2</code>, return <em>the minimum number of <strong>steps</strong> required to make</em> <code>word1</code> <em>and</em> <code>word2</code> <em>the same</em>.</p>

<p>In one <strong>step</strong>, you can delete exactly one character in either string.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> word1 = &quot;sea&quot;, word2 = &quot;eat&quot;
<strong>Output:</strong> 2
<strong>Explanation:</strong> You need one step to make &quot;sea&quot; to &quot;ea&quot; and another step to make &quot;eat&quot; to &quot;ea&quot;.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> word1 = &quot;leetcode&quot;, word2 = &quot;etco&quot;
<strong>Output:</strong> 4
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= word1.length, word2.length &lt;= 500</code></li>
	<li><code>word1</code> and <code>word2</code> consist of only lowercase English letters.</li>
</ul>

