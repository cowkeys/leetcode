package handle

// 自己的
/*
判断
先判断符号\空格
在判断数字
加入到一个数组
然后数组转化为数字
*/
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	nums := []int{}
	num := 0
	y := 1
	var f int
	for i := 0; i < len(str); i++ {
		if str[i] >= 48 && str[i] <= 57 {
			if len(nums) == 0 && str[i] == 48 {
				if f == 0 {
					f = 1
				}
				continue
			}
			nums = append(nums, int(str[i]-48))
			continue
		}

		if len(nums) == 0 {
			if str[i] == 45 && f == 0 {
				f = -1
				continue
			}

			if str[i] == 43 && f == 0 {
				f = 1
				continue
			}

			if str[i] != 32 {
				return 0
			}
			if f != 0 {
				return 0
			}
		}

		if len(nums) > 0 {
			break
		}
	}

	if len(nums) > 10 {
		if f >= 0 {
			return 2<<30 - 1
		}
		return -2 << 30
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > 0 {
			num += y * nums[i]
		}

		if f >= 0 && num > (2<<30-1) {
			return 2<<30 - 1
		}

		if num*f < -2<<30 {
			return -2 << 30
		}
		y *= 10

	}

	if f != 0 {
		num = num * f
	}
	return num
}

// 自动机
func myAutoAtoi(str string) int {

	MAX := 2<<30 - 1
	MIN := -2 << 30
	state := "start"
	table := make(map[string][]string, 0)
	table["start"] = []string{"start", "sign", "number", "end"}
	table["sign"] = []string{"end", "end", "number", "end"}
	table["number"] = []string{"end", "end", "number", "end"}
	table["end"] = []string{"end", "end", "end", "end"}

	var check = func(b byte) int {
		// 空格
		if b == 32 {
			return 0
		}
		// + / -
		if b == 45 || b == 43 {
			return 1
		}
		// 0~9
		if b >= 48 && b <= 57 {
			return 2
		}
		// else
		return 3
	}
	res := 0
	sign := 1
	for i := 0; i < len(str); i++ {
		state = table[state][check(str[i])]
		if state == "number" {
			res = res*10 + int(str[i]-48)
			//fmt.Println(res)
			if res > MAX && sign == 1 {
				return MAX
			}
			if sign == -1 && res > -MIN {
				return MIN
			}
		} else if state == "sign" {
			//fmt.Println("sign:", str[i], string(str[i]))
			if str[i] == 45 {
				sign = -1
			} else {
				sign = 1
			}
		}
	}

	return sign * res
}

func MyAtoi(str string) int {
	return myAutoAtoi(str)
}
