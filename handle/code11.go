package handle

//盛最多水的容器
// 暴力
func maxArea1(height []int) int {

	val := 0
	for i := 0; i < len(height); i++ {
		for j := 1; j < len(height); j++ {
			w := j - i
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			if w*h > val {
				val = w * h
			}
		}
	}

	return val
}

// 双指针
// 指向首尾, 数字小的往中间移动一格, 计算容量 如此反复
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	val := 0
	for i < j {
		w := j - i
		h := height[i]
		if h > height[j] {
			h = height[j]
			j--
		} else {
			i++
		}
		if w*h > val {
			val = w * h
		}
	}
	return val
}
