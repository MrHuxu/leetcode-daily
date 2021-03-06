\>\> [题目链接](https://leetcode.com/explore/challenge/card/august-leetcoding-challenge-2021/616/week-4-august-22nd-august-28th/3920/)

题意: 给定一个字符串 `preorder` 表示一棵二叉树的前序遍历序列, 其中空的树节点用 `#` 表示, 判断字符串是否是一个合法的前序遍历序列

解答: 栈模拟题, 每个叶子节点必然跟着两个 `#`, 那么遇到两个 `#` 的情况就弹三个元素出栈, 并压入一个 `#`, 最后栈里应该只剩一个 `#`

***original_content***

<p>One way to serialize a binary tree is to use <strong>preorder traversal</strong>. When we encounter a non-null node, we record the node&#39;s value. If it is a null node, we record using a sentinel value such as <code>&#39;#&#39;</code>.</p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/03/12/pre-tree.jpg" style="width: 362px; height: 293px;" />
<p>For example, the above binary tree can be serialized to the string <code>&quot;9,3,4,#,#,1,#,#,2,#,6,#,#&quot;</code>, where <code>&#39;#&#39;</code> represents a null node.</p>

<p>Given a string of comma-separated values <code>preorder</code>, return <code>true</code> if it is a correct preorder traversal serialization of a binary tree.</p>

<p>It is <strong>guaranteed</strong> that each comma-separated value in the string must be either an integer or a character <code>&#39;#&#39;</code> representing null pointer.</p>

<p>You may assume that the input format is always valid.</p>

<ul>
	<li>For example, it could never contain two consecutive commas, such as <code>&quot;1,,3&quot;</code>.</li>
</ul>

<p><strong>Note:&nbsp;</strong>You are not allowed to reconstruct the tree.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> preorder = "9,3,4,#,#,1,#,#,2,#,6,#,#"
<strong>Output:</strong> true
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> preorder = "1,#"
<strong>Output:</strong> false
</pre><p><strong>Example 3:</strong></p>
<pre><strong>Input:</strong> preorder = "9,#,#,1"
<strong>Output:</strong> false
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= preorder.length &lt;= 10<sup>4</sup></code></li>
	<li><code>preoder</code> consist of integers in the range <code>[0, 100]</code> and <code>&#39;#&#39;</code> separated by commas <code>&#39;,&#39;</code>.</li>
</ul>

