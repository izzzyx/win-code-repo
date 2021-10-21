package strandarray

// 剑指offer 题5 替换字符串中的空格
// 思路：
//	* 需要解决的问题：替换字符(%20)和被替换字符(" ")长度不同会覆盖数组中的其他字符
//		&空间复杂度（如果使用辅助数组而不是原地移动字符）
//		和时间复杂度（如果每次移动字符都移动空格后面所有的字符）
//	  ** 覆盖其他字符 => 先计算空格数量，预先往原字符串后面append替换需要的长度
//	  ** 空间&时间复杂度 => 不使用额外数组，原地移动字符，从最后一个字符开始挪
//	* 最终思路：用两个变量，一个记录原字符串尾部的字符位置（设为p2），一个记录替换字符串尾部的位置（设为p1），
//		然后就是挪的操作，也就是重复“把p1的值赋给p2，然后往前挪p1和p2”的操作，直到p1=p2
// 测试：
//	* 特殊输入：
//	  ** 0/nil值：空字符串；第一个字符是空字符；只有一个字符且是空字符
//	  ** 没有空字符
//  * 功能测试：
//	  ** 有连续多个空字符
// --------------------------------
// ---------- debug记录 -----------
//	* pos1是数组下标，但是赋值的时候没给它len-1
//	* 替换空格的的时候应该从p2往前倒推赋值，因为预留的位置都是在p1之后
//		** 并且p1只会往前移，从p1正着往后赋值占的是原字符串的位置
func ReplaceSpaceInString(s string) string {
	lenOriS := len(s)
	for i := 0; i < lenOriS; i++ {
		// s[i] is rune
		if string(s[i]) == " " {
			s += "  "
		}
	}

	pos2 := len(s) - 1
	// debug1
	// pos1 := lenOriS
	pos1 := lenOriS - 1
	b := []byte(s)
	replaceBytes := []byte("%20")
	spaceBytes := []byte(" ")
	for pos1 != pos2 {
		println(b[pos1])
		println(spaceBytes[0])
		if b[pos1] != spaceBytes[0] {
			b[pos2] = b[pos1]

			pos2--
		} else {
			// debug2
			// b[pos1] = replaceBytes[0]
			// b[pos1+1] = replaceBytes[1]
			// b[pos1+2] = replaceBytes[2]

			b[pos2-2] = replaceBytes[0]
			b[pos2-1] = replaceBytes[1]
			b[pos2] = replaceBytes[2]

			pos2 -= 3
		}

		pos1--
	}

	return string(b)
}
