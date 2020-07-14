package handle

// 罗马数字转数字
/*
自己的解法
从后往前
两个两个判断
若存在则num+=
若不存在则加单个数字的值
*/
func RomanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	/* R := make(map[string]int, 0)
	R["I"] = 1
	R["IV"] = 4
	R["V"] = 5
	R["IX"] = 9
	R["X"] = 10
	R["XL"] = 40
	R["L"] = 50
	R["XC"] = 90
	R["C"] = 100
	R["CD"] = 400
	R["D"] = 500
	R["CM"] = 900
	R["M"] = 1000 */
	var R = func(x string) int {
		switch x {
		case "I":
			return 1
		case "IV":
			return 4
		case "V":
			return 5
		case "IX":
			return 9
		case "X":
			return 10
		case "XL":
			return 40
		case "L":
			return 50
		case "XC":
			return 90
		case "C":
			return 100
		case "CD":
			return 400
		case "D":
			return 500
		case "CM":
			return 900
		case "M":
			return 1000
		}
		return 0
	}

	if len(s) == 1 {
		return R(s)
	}

	//"MCMXCIV"
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			num += R(string(s[i]))
			return num
		}

		b := string(s[i-1 : i+1])
		if v := R(b); v > 0 {
			num += v
			i--
			continue
		}

		num += R(string(s[i]))
	}
	return num
}

// 其他解法
/*

 */
func RomanToInt2(s string) int {
	if len(s) == 0 {
		return 0
	}

	var R = func(x string) int {
		switch x {
		case "I":
			return 1
		case "IV":
			return 4
		case "V":
			return 5
		case "IX":
			return 9
		case "X":
			return 10
		case "XL":
			return 40
		case "L":
			return 50
		case "XC":
			return 90
		case "C":
			return 100
		case "CD":
			return 400
		case "D":
			return 500
		case "CM":
			return 900
		case "M":
			return 1000
		}
		return 0
	}

	if len(s) == 1 {
		return R(s)
	}

	//"MCMXCIV"
	num := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := R(string(s[i]))
		if cur >= prev {
			num += cur
		} else {
			num -= cur
		}
		prev = cur
	}
	return num
}
