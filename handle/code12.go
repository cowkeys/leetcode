package handle

import "fmt"

/* 字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

// 枚举
func IntToRoman2(num int) string {
	R := make(map[int]string, 0)
	R[1] = "I"
	R[2] = "II"
	R[3] = "III"
	R[4] = "IV"
	R[5] = "V"
	R[6] = "VI"
	R[7] = "VII"
	R[8] = "VIII"
	R[9] = "IX"
	R[10] = "X"
	R[20] = "XX"
	R[30] = "XXX"
	R[40] = "XL"
	R[50] = "L"
	R[60] = "LX"
	R[70] = "LXX"
	R[80] = "LXXX"
	R[90] = "XC"
	R[100] = "C"
	R[200] = "CC"
	R[300] = "CCC"
	R[400] = "CD"
	R[500] = "D"
	R[600] = "DC"
	R[700] = "DCC"
	R[800] = "DCCC"
	R[900] = "CM"
	R[1000] = "M"
	R[2000] = "MM"
	R[3000] = "MMM"

	d := 1000
	x := 0
	res := ""

	if _, ok := R[num]; ok {
		return R[num]
	}
	for {
		x = num / d
		//fmt.Println("x", x, "d:", d)
		b := 0
		if x > 0 {
			res = res + R[x*d]
			b = x * d
		}
		num = num - b
		d /= 10
		if d == 0 {
			break
		}
	}

	fmt.Println(res)
	return res
}

// 贪心算法 最大开始
func IntToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	res := ""
	for i < 13 {
		for num >= nums[i] {
			res = res + roms[i]
			num -= nums[i]
		}
		i++
	}
	fmt.Println("res", res)
	return res
}
