package handle

// 自己的思路:
/*
将 数组一 [1,2,3,4,5] = 5 和数组二 [3,6,9]= 3 按顺序放入新的数组
[1,2,3,3..] 当新数组的长度是数组1和数组2 的一半的时候 停止放入
然后取 新数组位置在 (5+3)/2 的值, 注意要考虑两个数组长度的奇偶, 偶数就要两个数的平均值
时间复杂度应该是min(m,n)/2

*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	a := []int{}
	i, j := 0, 0
	var fn = func() []int {
		if i >= len1 {
			j++
			return []int{nums2[j-1]}
		}
		if j >= len2 {
			i++
			return []int{nums1[i-1]}
		}

		if nums1[i] < nums2[j] {
			i++
			return []int{nums1[i-1]}
		}

		if nums1[i] == nums2[j] {
			i++
			j++
			return []int{nums1[i-1], nums2[j-1]}
		}
		j++
		return []int{nums2[j-1]}

	}

	for i < len1 || j < len2 {
		a = append(a, fn()...)
		if len(a) >= (len1+len2)/2+1 {
			break
		}
	}

	if (len1+len2)%2 == 1 {
		return float64(a[(len1+len2)/2])
		//return float64(a[(len1+len2)/2])
	}

	return (float64(a[(len1+len2)/2-1]) + float64(a[(len1+len2)/2])) / 2
}
