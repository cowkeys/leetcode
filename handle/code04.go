package handle

import "fmt"

// 寻找两个正序数组的中位数

// 自己的思路:
/*
将 数组一 [1,2,3,4,5] = 5 和数组二 [3,6,9]= 3 按顺序放入新的数组
[1,2,3,3..] 当新数组的长度是数组1和数组2 的一半的时候 停止放入
然后取 新数组位置在 (5+3)/2 的值, 注意要考虑两个数组长度的奇偶, 偶数就要两个数的平均值
时间复杂度应该是min(m,n)/2

*/
func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
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

//  二分查找
/*
思路 将两个数组 分两组 左侧比右侧都小 分组后的两个数组 左边比右边都要小 并且数组1长度 == 数组2长度 或者 数组1长度 = 数组2长度+1 (奇偶)
如  [1,2,3,4,5] = 5 和数组二 [3,6,9]= 3  = 8 偶数
-> [1,2,3]|[4,5]
   [3]      |[6,9]  这样一来 中位数就是 左边最大值和右边最大值的平均数

如  [1,2,3,4] = 4 和数组二 [3,6,9]= 3  = 7 奇数
-> [1,2,3]|[4]
   [3]      |[6,9]  这样一来 中位数就是 左边最大值


*/

/*
 0 1   2 3
[1,4,| 6,8]
[2,5,|   9]

 0 1   2 3
[1,4,| 6,8]
[2,7,|   9]
*/
func FindMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	l1 := len(nums1)
	l2 := len(nums2)
	even := (l1+l2)%2 == 0

	//m := (l1 + l2 + 1) / 2

	var div = func(index int) int {
		if index/2 == 0 {
			return 1
		}
		return index / 2
	}
	/* m1 := nums1[(l1-1)/2]
	m2 := nums2[(l2-1)/2] */

	i, j := (l1-1)/2, (l2-1)/2
	if !even {
		i++
	}
	//k := 0
	var n1 = func(index int) int {
		if index >= len(nums1) {
			return nums1[len(nums1)-1]
		}
		return nums1[index]
	}

	var n2 = func(index int) int {
		if index >= len(nums2) {
			return nums2[len(nums2)-1]
		}
		return nums2[index]
	}

	for {
		/* 	if k > 10 {
			break
		}
		k++ */
		if j+1 >= l2 {
			j = l2 - 1
			break
		}
		if i+1 >= l1 {
			i = l1 - 1
			break
		}
		fmt.Println(nums1[i], nums2[j+1], "&&", nums2[j], nums1[i+1], nums1[i] < nums2[j+1] && nums2[j] < nums1[i+1])
		if nums1[i] < nums2[j+1] && nums2[j] < nums1[i+1] {
			break
		}
		if nums2[j] > nums1[i+1] {
			j -= div(j) //j / 2
			i += div(i) //i / 2
		}
		if nums1[i] > nums2[j+1] {
			i -= div(i) //i / 2
			j += div(j) //j / 2
		}
	}
	fmt.Println("i", i, "j", j)

	if even {
		if i+1 == len(nums1) {

		}
		fmt.Println("max", n1(i), n2(j), "max(", n1(i+1), n2(j+1), ")", "avg")
		return 0
	}
	fmt.Println(nums1[i], nums2[j], "max")

	// 边界
	/* if nums1[0] >= nums2[l2-1] {
		if even {

			return nums2[n2]
		}
		if l1 < m {
			return nums2[]
		}
	} */
	return 0
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 长度
	l1 := len(nums1)
	l2 := len(nums2)

	// 奇偶
	//l1odd := l1%2 == 1
	//l2odd := l2%2 == 1

	// 整体奇偶
	odd := (l1+l2)%2 == 1

	// 游标
	c1 := l1 / 2
	c2 := l2 / 2

	if 2*c1+2*c2-l1-l2+4 > 1 {
		c2--
	}
	/* if c1+c2+2-(l1-c1-1+l2-c2-1) > 1 {
		c2--
	} */
	/* if !l1odd {
		c1--
	}
	if !l2odd {
		c2--
	} */
	fmt.Println(nums1, nums2)
	fmt.Printf("c1:%d,c2:%d\n", c1, c2)
	for nums1[c1] > nums2[c2+1] || nums2[c2] > nums1[c1+1] {
		fmt.Printf("before c1:%d,c2:%d\n", c1, c2)
		if nums1[c1] > nums2[c2+1] {
			fmt.Println("nums1[c1] > nums2[c2+1]", nums1[c1], nums2[c2+1])
			if c1 <= 2 {
				c1--
			} else {
				c1 -= c1 / 2
			}

			if l2-c2-1 <= 2 {
				c2++
			} else {
				c2 += (l2 - c2) / 2
			}
		} else if nums2[c2] > nums1[c1+1] {
			fmt.Println("nums2[c2] > nums1[c1+1]", nums2[c2], nums1[c1+1])
			if c2 <= 2 {
				c2--
			} else {
				c2 -= c2 / 2
			}

			if l1-c1-1 <= 2 {
				c1++
			} else {
				c1 += (l1 - c1) / 2
			}
		}

		fmt.Printf("after c1:%d,c2:%d\n", c1, c2)
		if c1 < 0 || c2 >= l2-1 {
			break
		}
		if c2 < 0 || c1 >= l1-1 {
			break
		}
	}

	leftmax := 0
	if c1 < 0 {
		leftmax = nums2[c2]
	} else if c2 < 0 {
		leftmax = nums1[c1]
	} else if nums1[c1] >= nums2[c2] {
		leftmax = nums1[c1]
	} else if nums2[c2] >= nums1[c1] {
		leftmax = nums2[c2]
	}

	rightmin := 0
	if c1 >= l1-1 {
		rightmin = nums2[c2+1]
	} else if c2 >= l2-1 {
		rightmin = nums1[c1+1]
	} else if nums1[c1+1] <= nums2[c2+1] {
		leftmax = nums1[c1+1]
	} else if nums2[c2+1] <= nums1[c1+1] {
		leftmax = nums2[c2+1]
	}

	fmt.Println("c1,", c1, "c2", c2)
	fmt.Println("leftmax", leftmax, "rightmin", rightmin)
	if odd {
		return float64(leftmax)
	}

	return float64(float64(leftmax+rightmin) / 2)
}

// 1,4,7,10,29 -- 4
// 0, 4
//
func TwoFind(list []int, v int) int {
	fmt.Println(list)
	left := 0
	right := len(list) - 1

	mid := -1
	for left <= right {
		mid = (right + left) / 2
		fmt.Println("left,", left, "right", right, "mid", mid)
		if list[mid] == v {
			break
		}
		if list[mid] < v {
			left = mid
			continue
		}
		right = mid
	}
	fmt.Println("mid", mid)
	return mid
}

func BinarySearch(n []int, target int) int {
	length := len(n)
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if n[mid] > target {
			high = mid - 1
		} else if n[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
