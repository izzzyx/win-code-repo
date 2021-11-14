package hashusedsearch

// 剑指offer 题50 找出数组中第一个只出现一次的字符，如果没有则返回一个单空格
// 思路：利用hash表存储字符出现次数；进一步优化空间复杂度可以自己实现hash表（数组实现），用ascii码做key
// 功能测试：能找到target字符
// 边界测试：
//	*0值/nil值：字符数组长度为0/nil；不存在满足条件的字符（所有字符都出现了多次）
//	*1值：字符数组长度为1；所有字符都只出现了一次
// 麻烦的点在于理解go里面的字符串、[]byte和rune的关系. ToT
func FirstUniqChar(s string) byte {
	if len(s) == 0 {
		// 显式转一下
		return []byte(" ")[0]
	}
	if len(s) == 1 {
		return s[0]
	}

	charToCounterMap := make(map[byte]int)
	// go中range map并不会按插入顺序
	charToIndexMap := make(map[byte]int)
	for i := range s {
		//debug1:没考虑删除element之后，index也变了的情况，需要更新charToIndexMap中存储的index
		// if indexInUniqArr, ok := charToIndexMap[s[i]]; ok {
		// 	uniqArr = append(uniqArr[:indexInUniqArr], uniqArr[indexInUniqArr+1:]...)
		// } else {
		// 	uniqArr = append(uniqArr, s[i])
		// 	charToIndexMap[s[i]] = len(uniqArr) - 1
		// }
		num, ok := charToCounterMap[s[i]]
		if ok {
			charToCounterMap[s[i]] = num + 1
		} else {
			charToCounterMap[s[i]] = 1
			charToIndexMap[s[i]] = i
		}
	}

	// debug2:比较的时候怎么会给初始值赋值0呢？？
	// targetIndex := 0
	targetIndex := len(s) - 1
	hits := false
	for k, v := range charToCounterMap {
		if v == 1 {
			hits = true
			if charToIndexMap[k] < targetIndex {
				targetIndex = charToIndexMap[k]
			}
		}
	}

	if hits {
		return s[targetIndex]
	} else {
		return " "[0]
	}
}
