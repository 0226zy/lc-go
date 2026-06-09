package sorts

// MergeSort 归并排序（升序）
// 对切片 nums 进行原地升序排序。
// 时间复杂度: O(n log n)  空间复杂度: O(n)
func MergeSort(nums []int) {
	mergeSort(nums, func(leftVal, rightVal int) bool {
		return leftVal <= rightVal
	})
}

// MergeSortDesc 归并排序（降序）
// 对切片 nums 进行原地降序排序。
// 时间复杂度: O(n log n)  空间复杂度: O(n)
func MergeSortDesc(nums []int) {
	mergeSort(nums, func(leftVal, rightVal int) bool {
		return leftVal >= rightVal
	})
}

func mergeSort(nums []int, takeLeft func(leftVal, rightVal int) bool) {
	if len(nums) <= 1 {
		return
	}
	tmp := make([]int, len(nums))
	mergeHelper(nums, tmp, 0, len(nums)-1, takeLeft)
}

func mergeHelper(
	nums, tmp []int,
	left, right int,
	takeLeft func(leftVal, rightVal int) bool,
) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeHelper(nums, tmp, left, mid, takeLeft)
	mergeHelper(nums, tmp, mid+1, right, takeLeft)

	// 两个区间首尾已经满足整体有序时，直接跳过本次合并。
	if takeLeft(nums[mid], nums[mid+1]) {
		return
	}

	merge(nums, tmp, left, mid, right, takeLeft)
}

func merge(
	nums, tmp []int,
	left, mid, right int,
	takeLeft func(leftVal, rightVal int) bool,
) {
	i,j,k:=left,mid+1,left
	for i<=mid && j<=right{
		if takeLeft(nums[i],nums[j]){
			tmp[k]=nums[i]
			i++
		}else{
			tmp[k]=nums[j]
			j++
		}
		k++
	}
	for i<=mid{
		tmp[k]=nums[i]
		i++
		k++
	}
	for j<=right{
		tmp[k]=nums[j]
		j++
		k++
	}
	copy(nums[left:right+1],tmp[left:right+1])
}
