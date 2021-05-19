\>\> [题目链接](https://leetcode.com/explore/featured/card/may-leetcoding-challenge-2021/600/week-3-may-15th-may-21st/3745/)

题意: 给定一棵二叉树, 假设我们可以在二叉树的任意节点上防止相机, 每个相机可以拍到当前节点, 父节点和左右子节点, 返回拍到所有二叉树所有节点需要的最少的相机数  
  
解答: 使用 dfs 的方式求解, 遍历的过程中, 每个节点可能出现三个状态  
  
1. `NULL`: 当前节点未被相机覆盖  
2. `COVER`: 当前节点已被相机覆盖  
3. `CAMERA`: 当前节点放置相机  
  
每次遍历都返回当前节点的状态以及当前节点为根节点的树存在多少个相机, 那么在分别得到左右子节点状态以及左右子树相机数后, 有如下几个判断分支  
  
1. 如果左右子节点任意一个为 `NULL`, 那么需要在当前节点设置相机去覆盖子节点, 并且总相机数需加 1  
2. 如果左右子节点任意一个为 `CAMERA`, 那么当前节点可被子节点覆盖, 无须增加相机数  
3. 如果不符合 1 和 2, 说明两个子节点全都被覆盖了, 当前节点返回 `NULL`, 通过后续的遍历被父节点覆盖  
  
额外有一个情况, 就是 dfs 整体返回的状态是 NULL, 即根节点未被覆盖, 则需要对结果加 1 即在根节点上放置相机
  
最后返回即可

***original_content***

<p>Given a binary tree, we install cameras on the nodes of the tree.&nbsp;</p>

<p>Each camera at&nbsp;a node can monitor <strong>its parent, itself, and its immediate children</strong>.</p>

<p>Calculate the minimum number of cameras needed to monitor all nodes of the tree.</p>

<p>&nbsp;</p>

<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_01.png" style="width: 138px; height: 163px;" />
<div>
<pre>
<strong>Input: </strong><span id="example-input-1-1">[0,0,null,0,0]</span>
<strong>Output: </strong><span id="example-output-1">1</span>
<strong>Explanation: </strong>One camera is enough to monitor all nodes if placed as shown.
</pre>

<div>
<p><strong>Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_02.png" style="width: 139px; height: 312px;" />
<pre>
<strong>Input: </strong><span id="example-input-2-1">[0,0,null,0,null,0,null,null,0]</span>
<strong>Output: </strong><span id="example-output-2">2
<strong>Explanation:</strong> At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.</span>
</pre>

<p><br />
<strong>Note:</strong></p>

<ol>
	<li>The number of nodes in the given tree will be in the range&nbsp;<code>[1, 1000]</code>.</li>
	<li><strong>Every</strong> node has value 0.</li>
</ol>
</div>
</div>

