package search

// 剑指 题3 找出数组中的任一重复数字
// 思路：根据要查找的数组的特征（长度为n，数值范围在0~n-1的范围内）
//	就可以利用数组下标的数组值的关系，把数值都放到与它相等的下标，
//	总有一个下标位置挤不下（这个下标值重复了）
//  完全是脑筋急转弯啊
// 测试：
//	* 功能测试：
//	  * 正常带有一个重复数字的nums
//	  * 多个不同的重复数字
//	* 边界测试：
//	  * 0/nil值：nums长度为0；
//	* 特殊输入：
//	  * nums没有重复数字
//	* 错误输入：
//	  * nums中的数据范围不在0~n-1
func FindRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums) && nums[i] < len(nums); {
		if nums[i] == i {
			// nums[i]的值放对了位置
			i++
			continue
		}

		m := nums[i]
		if nums[m] == m {
			return m
		} else {
			// 交换index i和index m的值，把值m放到index m的位置上
			nums[i] = nums[m]
			nums[m] = m
		}
	}

	// no repeat num
	return -1
}
