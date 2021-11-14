package search

// 剑指offer 题53-Ⅱ 找出长度为n-1的数组中缺失的数字。
// 数组特征为：1.递增排序数组 2.取值范围为[0, n-1]且没有重复数字
// 思路：找缺失的数字 => 怎么找 => 缺失的数字有什么特点？=> 举例画图
// 测试：
//	*功能测试：缺失的数字位于数组开头、中间、末尾;位于二分查找的左边/右边/刚好
//  *边界测试：数组长度为0；数组只有一个数字(n-1=1)；
//	*特殊输入：数组没有缺失数字
func MissingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	// 缺失的数字在开头
	if nums[0] != 0 {
		return 0
	}
	// 缺失的数字在末尾
	if nums[len(nums)-1] == len(nums)-1 {
		return len(nums)
	}
	// 缺失的数字在中间，查找m：nums[m]≠m，且nums[m-1]=m-1
	// debug2: end := len(nums)-1 不用减1，否则要考奇偶的情况
	end := len(nums)
	for m := len(nums) / 2; m > 0 && m < len(nums); {
		if nums[m] != m {
			if m-1 >= 0 && nums[m-1] == m-1 {
				return m
			} else {
				// debug1: 没考虑else的情况
				// 找左边
				end = m
				m = m / 2
			}
		} else {
			// 找右边
			m = (end-m)/2 + m
		}
	}

	return -1
}
