package handle

// 最长回文子串 "babad" > bab
/*
自己的思路 反向扩散 从外围开始
babad
先查看 s[0] == s[len] 即 b == d (不相等就break)
若相等继续 看s[1] == s[len-1] (a==a)...
若一直往下直到汇合都相等 则储存res

若有一个不相等则退出
继续判断 从 s[0] == s[len-1]
		  s[1] == s[len-2]...

第二轮 从s[1] == s[len]
		s[2] == s[len-1]...
若res的长度 > 之前的 则替换res
返回 res

其他思路:
1 中心扩散 选择一个数,从两边扩散如果一直相等则是回文,记录长度
2 动态规划 画一个矩阵图,
*/
func LongestPalindrome(s string) string {
	var check = func(item string) bool {
		return item[0] == item[len(item)-1]
	}
	res := ""
	for i := 0; i < len(s); i++ {
		index := i
		j := len(s)
		if j-index <= len(res) {
			break
		}
		for j > i {
			//j = len(s)
			str := s[i:j]
			bk := str
			brk := false
			index = 0
			for check(str) {
				index++
				if index >= j-index-i {
					if len(bk) > len(res) {
						res = bk
					}
					brk = true
					break
				}

				str = bk[index : j-index-i]
			}
			if brk {
				break
			}
			j--
		}
	}
	return res
}
