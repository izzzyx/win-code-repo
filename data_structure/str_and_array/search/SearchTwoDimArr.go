package search

// 剑指offer 题4：二维数组中的查找，在一个横向递增&竖向递增的二维数组中查找是否包含指定整数
// 类型tag：查找算法；数据结构（数组）
// 解题思路：
//		1. 从右上角开始查询
//		2. 根据二维数组递增的特征，每次查找都能剔除一行或者一列
//		3. 查找余下的区域
// 测试：
// 	* 功能测试：
//		* 正常的数组&要查找的数存在，return true
//		* 要查找的数字不存在于数组中, return false
//	* 边界测试：
//		* 二维数组长度为0
//		* 二维数组只有行/只有列；二维数组只有一个元素；
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	// 从右上角开始查找
		
	i := 0
	j := len(matrix[0]) - 1
	var find bool
	for j >= 0 && i < len(matrix) {
		if find {
			return true
		}
		// 比较current和target
		if matrix[i][j] > target {
			// 剔除current所在的这一列，移动i,j指针
			j--
		} else if matrix[i][j] < target {
			// 剔除current所在的这一行，移动i,j指针
			i++
		} else {
			find = true
		}
	}

	return find
}
