package rotateandsearch

// 剑指 Offer 11. 旋转数组的最小数字:输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素
// 思路：
//	1. 原始数组的特征：递增 => 旋转后的数组特征：以旋转数组为界，
//		划分为左右两个递增的子数组，且左子数组的所有数字都大于右子数组
//	2. 利用二分查找找到分界线：左右子树有序&分界线的特点
// 测试：
//	*功能测试：正常数组，return最小的数
//	*边界测试：
//		*0值/nil值：数组长度为0/nil
//		*1值：数组长度为1
//		*数组头尾相等：[3,1,3,3]
//	*特殊输入：
//		*分界线在数组末尾：[2,3,4,5,1]
//		*完全不旋转的：比如[1,3,5]
//		*数组中存在相等的值：[2,2,0,1]
func MinArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	// 二分查找
	mid := len(numbers) / 2
	end := len(numbers)
	// 用于确定mid当前是在左子数组还是右子数组
	first := numbers[0]
	last := numbers[len(numbers)-1]
	for mid != end {
		hit := false
		if numbers[mid] < numbers[mid-1] &&
			(mid+1 == len(numbers) || numbers[mid] <= numbers[mid+1]) {
			return numbers[mid]
		}
		// 查左半边: current > current - 1 && current < first
		// debug1：没考虑数组中存在相等的元素（不是递增数组吗？？哪来的相等元素！）
		if numbers[mid] >= numbers[mid-1] && numbers[mid] < first {
			hit = true
			end = mid + 1
			mid = end / 2
		}
		// 查右半边：current > current - 1 && current > last
		if numbers[mid] >= numbers[mid-1] && numbers[mid] > last {
			hit = true
			mid = (end-mid)/2 + mid
		}
		// 找不出原始数组
		// debug2: 找不出原始数组的情况只考虑了完全不旋转，没有考虑first==last以至于分不出左右子树的情况
		if !hit {
			// 完全不旋转的情况，比如[1,2,3,4,5]
			if first < last {
				return first
			} else {
				// first==last以至于分不出左右子树的情况，顺序查找
				minNum := numbers[0]
				for i := 0; i < end-1; i++ {
					if numbers[i] < minNum {
						minNum = numbers[i]
					}
				}
				return minNum
			}
		}
	}

	return -1
}
