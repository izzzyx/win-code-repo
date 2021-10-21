package strandarray

// 剑指offer 题58 将一个字符串反转，但保持单词，比如I am a student翻转为student a am I
// 思路：翻转两次；先翻转整个字符串，再翻转字符串中的每个单词
//	* I am a student -> tneduts a ma I -> student a am I
//	* 因此只需要实现一个完全翻转字符串的函数，第一次传整个字符串，第二次传其中的单词
// 剑指offer 题58 Ⅱ 左旋字符串；给定参数position和字符串，将字符串沿position左旋
// 思路；沿pos将字符串分为左右两部分（index从0开始，pos属于右半部分）&分别翻转两部分&翻转整个字符串
//	* abcdefg -> bagfedc -> cdefgab
// （脑筋急转弯的范畴了属于是）
// 测试：
//	* 功能测试：输入一个正常字符串和位于其中的position
//	* 特殊输入：
//	  ** nil/0：字符串长度为0；position位置为0/结尾
//	  ** 1：字符串长度为1，position也=1
//	* 错误输入：
//	  ** position越界/<0
func ReverseLeftWords(s string, n int) string {
	if len(s) == 0 || n > len(s) || n < 0{
		return ""
	}

	leftStr := s[:n]
	rightStr := s[n:]
	
	rLeft := ReverseStr(leftStr)
	rRight := ReverseStr(rightStr)

	return ReverseStr(rLeft+rRight)
}

func ReverseStr(s string) string {
	if len(s) == 0 {
		return ""
	}
	// 用两个指针p1p2，向中间移动并一直交换值直到相等（奇数）或p2+1=p1（偶数）
	pos1 := 0
	pos2 := len(s) - 1
	b := []byte(s)
	for {
		if pos1==pos2 || pos2+1==pos1 {
			break
		}
		
		tmp := b[pos1]
		b[pos1] = b[pos2]
		b[pos2] = tmp

		pos1++
		pos2--
	}

	return string(b)
}
