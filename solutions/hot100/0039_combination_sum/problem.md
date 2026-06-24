# 39. 组合总和 (Combination Sum)

## 题目描述
给你一个无重复元素的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的所有不同组合，并以列表形式返回。你可以按任意顺序返回这些组合。

candidates 中的同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。

## 示例
```
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]

输入：candidates = [2,3,5], target = 8
输出：[[2,2,2,2],[2,3,3],[3,5]]
```

## 约束
- 1 <= candidates.length <= 30
- 2 <= candidates[i] <= 40
- candidates 的所有元素互不相同
- 1 <= target <= 40

## 分类
回溯
