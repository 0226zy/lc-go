package uniquebinarysearchtrees

// NumTrees 不同的二叉搜索树（标准 DP 数组版）
// 求由 n 个节点（值 1..n）能组成多少种结构不同的二叉搜索树。
// dp[i] 表示 i 个节点能组成的 BST 种数，dp[i] = Σ dp[j-1] * dp[i-j]（j 从 1 到 i，即以第 j 个值为根）。
// 时间复杂度: O(n²)  空间复杂度: O(n)
func NumTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1 // base case：空树和单节点树都只有 1 种
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 以第 j 个值为根：左子树 j-1 个节点，右子树 i-j 个节点
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// NumTreesAlternative 不同的二叉搜索树（卡特兰数公式版）
// 答案即第 n 个卡特兰数，递推：C(n+1) = C(n) * 2 * (2n+1) / (n+2)。
// 每项只依赖前一项，用单个变量迭代即可；必须先乘后除保证整除。
// 时间复杂度: O(n)  空间复杂度: O(1)
func NumTreesAlternative(n int) int {
	c := 1 // C(0) = 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}
