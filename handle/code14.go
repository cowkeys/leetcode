package handle

import "fmt"

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := 0
	s := ""
Loop:
	for {
		var c byte
		for k := 0; k < len(strs); k++ {
			if i > len(strs[k])-1 {
				break Loop
			}
			if c == 0 {
				c = strs[k][i]
			}
			if strs[k][i] != c {
				break Loop
			}
		}
		fmt.Println("--", string(c))
		s += string(c)
		i++
	}
	return s
}
