package handle

// abcabcbb 自己的做的
func LengthOfLongestSubstringOld(s string) int {
	m := make(map[byte]bool, 0)
	max := 0
	for i, j := 0, 0; i < len(s); i++ {
		if !m[s[i]] {
			m[s[i]] = true
			continue
		}
		if len(m) > max {
			max = len(m)
		}
		m = make(map[byte]bool, 0)
		j++
		i = j - 1
	}
	if len(m) > max {
		return len(m)
	}
	return max
}

// abcabcbb 自己的做的2 使用滑动窗口
func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool, 0)
	max := 0
	i, j := 0, 0
	for i < len(s) {
		for m[s[i]] {
			delete(m, s[j])
			j++
		}
		m[s[i]] = true
		i++
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}

// 别人的思路
/* 思路：
这道题主要用到思路是：滑动窗口

什么是滑动窗口？

其实就是一个队列,比如例题中的 abcabcbb，进入这个队列（窗口）为 abc 满足题目要求，当再进入 a，队列变成了 abca，这时候不满足要求。所以，我们要移动这个队列！

如何移动？

我们只要把队列的左边的元素移出就行了，直到满足题目要求！

一直维持这样的队列，找出队列出现最长的长度时候，求出解！

作者：powcai
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 */
