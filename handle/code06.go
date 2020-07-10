package handle

// Z 字形变换
/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

1 L     D     R
2 E   O E   I I
3 E C   I H   N
4 T     S     G

1 L     I
2 E   E S   G
3 E  D  H  N
4 T O   I I
5 C     R

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Convert(s string, numRows int) string {
	return convert(s, numRows)
}

// https://leetcode-cn.com/problems/zigzag-conversion/solution/z-zi-xing-bian-huan-by-leetcode/
// 按行访问
/*
执行用时：4 ms, 在所有 Go 提交中击败了94.68%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func convert(s string, numRows int) string {
	var res []byte
	lenIndex := len(s) - 1
	if lenIndex <= numRows-1 {
		return s
	}
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		lineIndex := i
		index := lineIndex
		p := numRows + numRows - 2
		skip := 2 * (numRows - lineIndex - 1)
		k := skip

		for {
			if index > lenIndex {
				break
			}
			res = append(res, s[index])
			if k == 0 {
				index += p - k
			}
			index += k
			if p-k > 0 {
				k = p - k
			}
		}
	}
	return string(res)
}
