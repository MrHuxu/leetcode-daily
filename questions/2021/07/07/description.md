\>\> [题目链接](https://leetcode.com/explore/featured/card/july-leetcoding-challenge-2021/608/week-1-july-1st-july-7th/3805/)

题意: 给定一个 n x n 的矩阵, 每一行每一列都递增, 找出里面第 k 个大的元素

解答: 比较简单粗暴的做法, 声明一个 map, 首先最小元素肯定在 [0,0] 的位置, 那么下一个元素肯定是在 [0,0] 的右边或下边, 找到这两个位置里较小的那个, 将索引放进 map, 那么每次找下一个元素, 其实就是找数组里已有节点的右边或下边节点里最小的那个, 同时如果一个节点的右边和下边节点都用过, 这个节点就不需要遍历了, 可以从 map 中删除, 最后返回 map 中 k 对应的坐标的值即可

***original_content***

<p>Given an <code>n x n</code> <code>matrix</code> where each of the rows and columns are sorted in ascending order, return <em>the</em> <code>k<sup>th</sup></code> <em>smallest element in the matrix</em>.</p>

<p>Note that it is the <code>k<sup>th</sup></code> smallest element <strong>in the sorted order</strong>, not the <code>k<sup>th</sup></code> <strong>distinct</strong> element.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
<strong>Output:</strong> 13
<strong>Explanation:</strong> The elements in the matrix are [1,5,9,10,11,12,13,<u><strong>13</strong></u>,15], and the 8<sup>th</sup> smallest number is 13
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> matrix = [[-5]], k = 1
<strong>Output:</strong> -5
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>n == matrix.length</code></li>
	<li><code>n == matrix[i].length</code></li>
	<li><code>1 &lt;= n &lt;= 300</code></li>
	<li><code>-10<sup>9</sup> &lt;= matrix[i][j] &lt;= 10<sup>9</sup></code></li>
	<li>All the rows and columns of <code>matrix</code> are <strong>guaranteed</strong> to be sorted in <strong>non-decreasing order</strong>.</li>
	<li><code>1 &lt;= k &lt;= n<sup>2</sup></code></li>
</ul>

