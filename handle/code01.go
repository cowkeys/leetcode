package handle

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k := 0; k < len(nums); k++ {
		if _, ok := m[target-nums[k]]; ok {
			return []int{m[target-nums[k]], k}
		}
		m[nums[k]] = k
	}

	return []int{}
}
