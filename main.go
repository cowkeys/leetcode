package main

import (
	"fmt"
	"os"

	"github.com/cowkeys/leetcode/handle"
)

func main() {
	//fmt.Println(handle.TwoSum([]int{2, 7, 11, 15}, 9))
	switch os.Args[1] {
	case "1":
		fmt.Println(handle.TwoSum([]int{3, 3}, 6))
	case "2":
		/* l1 := parseListNode([]int{2, 4, 3})
		l2 := parseListNode([]int{5, 6, 4}) */

		/* l1 := parseListNode([]int{1, 8})
		l2 := parseListNode([]int{0})
		*/
		/* l1 := parseListNode([]int{9, 8})
		l2 := parseListNode([]int{1}) */

		/* l1 := parseListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
		l2 := parseListNode([]int{5, 6, 4}) */

		/* l1 := parseListNode([]int{5})
		l2 := parseListNode([]int{5}) */

		l1 := handle.ParseListNode([]int{9})
		l2 := handle.ParseListNode([]int{9})

		res := handle.AddTwoNumbers(l1, l2)

		for {
			fmt.Printf("%d", res.Val)
			res = res.Next
			if res == nil {
				break
			}
		}
	case "3":
		fmt.Println(handle.LengthOfLongestSubstring("abcabcbb") == 3)
		fmt.Println(handle.LengthOfLongestSubstring("bbbbb") == 1)
		fmt.Println(handle.LengthOfLongestSubstring("pwwkew") == 3)
		fmt.Println(handle.LengthOfLongestSubstring("c") == 1)
		fmt.Println(handle.LengthOfLongestSubstring("") == 0)
		fmt.Println(handle.LengthOfLongestSubstring(" ") == 1)
		fmt.Println(handle.LengthOfLongestSubstring("aab") == 2)
		fmt.Println(handle.LengthOfLongestSubstring("dvdf") == 3)
		fmt.Println(handle.LengthOfLongestSubstring("qrsvbspk") == 5)

		//"abcabcbb"
	case "4":
		fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 100, 200, 300, 400, 500}) == 52.5)
		fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}) == 3.5)

		fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{2, 4}) == 3)

		fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{100, 200}) == 5)

		fmt.Println(handle.FindMedianSortedArrays([]int{1, 2}, []int{1, 2}) == 1.5)
	case "5":

		/* fmt.Println(handle.LongestPalindrome("bbbbb") == "bbbbb", "bbbbb")

		fmt.Println(handle.LongestPalindrome("abcda") == "a", "abcda")
		fmt.Println(handle.LongestPalindrome("babad") == "bab", "babad")

		fmt.Println(handle.LongestPalindrome("cbbd") == "bb", "cbbd")

		fmt.Println(handle.LongestPalindrome("abacab") == "bacab", "abacad")
		fmt.Println(handle.LongestPalindrome("aba") == "aba", "aba")
		*/
		//fmt.Println(handle.LongestPalindrome("abcba") == "abcba", "abcba")

		var fn = func(s, as string) {
			fmt.Println("test", s, handle.LongestPalindrome(s) == as)
		}

		fn("bbbbb", "bbbbb")
		fn("abcda", "a")
		fn("babad", "bab")
		fn("cbbd", "bb")
		fn("abacab", "bacab")
		fn("aba", "aba")

		fn("abcba", "abcba")
		fn("babadada", "adada")
	}

}
