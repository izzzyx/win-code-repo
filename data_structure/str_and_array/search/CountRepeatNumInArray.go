package search

// 剑指offer 题53 计算指定数字在排序数组中出现的次数
// 思路：简单可以想到直接遍历计数；复杂的则考虑用二分查找（除2，除2，除2……，花费时间相当于2的n次方分之一即log2n）
// 	 *利用给定的数组的特性（有序）&二分查找。有序数组=>找到第一个n和最后一个n，它们之间的即是所有想找的数字（闭区间）
// 测试：
//	*功能测试：正常的有序nums和target；
//	*边界测试：
//	  *nil/0值（nums长度为0，target不在数组中）
//	  *1值：target在nums的第一个且仅有这一个；target在nums的最后一个且仅有这一个;
//	*错误输入：nums无序；
func CountRepeatNumInArray(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	// getFirstK&lastK
	firstK := -1
	lastK := -1
	a := 0
	b := len(nums)
	// debug5: a <= b
	for mid := 0; a < b; {
		// mid = 对半分 中线左边的数的下标
		// 偶数
		// debug2:
		// currentArr = nums[:mid] or currentArr = nums[mid+1:]
		// mid = len(currenArr) / 2
		// debug3:
		// b-a == 0
		if (b-a)%2 == 0 { // 最小都是2
			mid = (b-a)/2 + a - 1
		} else {
			// 奇数
			mid = (b-a)/2 + a
		}

		// debug1：mid == target
		if nums[mid] == target {
			if mid-1 < 0 {
				firstK = mid
				break
			}
			b = mid // firstK
		}
		if nums[mid] < target {
			// debug4: if mid+1 > len(nums)
			if mid+1 >= len(nums) {
				// target不存在于nums中
				return 0
			}
			if nums[mid+1] == target {
				firstK = mid + 1
				break
			}
			a = mid + 1
		}
		if nums[mid] > target {
			// debug6: a = mid + 1
			b = mid
		}
	}

	a = 0
	b = len(nums)
	for mid := 0; a < b; {
		// mid = 对半分 中线左边的数的下标
		// 偶数
		if (b-a)%2 == 0 { // 最小都是2
			mid = (b-a)/2 + a - 1
		} else {
			// 奇数
			mid = (b-a)/2 + a
		}

		if nums[mid] > target {
			if mid-1 < 0 {
				// target不存在于nums中
				return 0
			}
			if nums[mid-1] == target {
				lastK = mid - 1
				break
			}
			b = mid
		}
		if nums[mid] == target {
			if mid+1 >= len(nums) {
				lastK = mid
				break
			}
			a = mid + 1 // lastK
		}
		if nums[mid] < target {
			a = mid + 1
		}
	}

	// 没找到
	if lastK == -1 || firstK == -1 {
		return 0
	}
	return lastK - firstK + 1
}
