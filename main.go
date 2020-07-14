package main

import (
	"fmt"
	"os"
	"time"

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
	case "44":
		handle.TwoFind([]int{1, 3, 4, 8, 10, 40, 48, 60, 80, 100}, 8)
	case "4":
		var fn = func(x, y []int, as float64) {
			res := handle.FindMedianSortedArrays(x, y)
			fmt.Println(res, as, res == as)
		}
		fn([]int{7, 8, 9, 10}, []int{1, 2, 3, 4, 5}, 5)
		fn([]int{1, 3, 5}, []int{2, 4, 100, 200, 300, 400, 500}, 52.5)
		/* fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 100, 200, 300, 400, 500}) == 52.5)
		fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}) == 3.5) */
		//fmt.Println(handle.FindMedianSortedArrays([]int{7, 8, 9, 10}, []int{1, 2, 3, 4, 5}) == 5)
		/*
			fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{2, 4}) == 3)

			fmt.Println(handle.FindMedianSortedArrays([]int{1, 3, 5}, []int{100, 200}) == 5)

			fmt.Println(handle.FindMedianSortedArrays([]int{1, 2}, []int{1, 2}) == 1.5) */
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
			handle.LongestPalindrome(s)
			fmt.Println("test", s, handle.LongestPalindrome(s) == as)
		}
		now := time.Now().UnixNano()
		for i := 0; i < 1; i++ {
			fn("bbbbb", "bbbbb")
			fn("abcda", "a")
			fn("babad", "bab")
			fn("cbbd", "bb")
			fn("abacab", "bacab")
			fn("aba", "aba")
			fn("abcba", "abcba")
			fn("babadada", "adada")
		}

		fmt.Println("cost", (time.Now().UnixNano()-now)/1e6, "ms", time.Now().UnixNano())
	case "6":
		var fn = func(s string, n int, as string) {
			res := handle.Convert(s, n)

			fmt.Println(res, as, res == as)
		}

		fn("LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN")
		fn("LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG")
		fn("AB", 1, "AB")
	case "7":
		var fn = func(x int) {
			handle.Reverse(x)
			fmt.Println(handle.Reverse(x))
		}
		//fn(123)
		fn(-1234)
	case "8":
		var fn = func(x string, as int) {
			res := handle.MyAtoi(x)
			fmt.Println("result:", x, res, res == as)
		}
		//fn(123)
		fn("-42", -42)
		fn("a234", 0)
		fn("4193 with words", 4193)
		fn("words and 987", 0)
		fn(" -43", -43)
		fn(" 43", 43)
		fn("9223372036854775808", 2147483647)
		fn("-   234", 0)
		fn("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459", 2147483647)
		//fn("100001", 100001)
		fn("  0000000000012345678", 12345678)
	case "9":
		var fn = func(x int, as bool) {
			res := handle.IsPalindrome(x)
			fmt.Println(res == as)
		}
		fn(1234321, true)
		fn(1234, false)
		fn(10, false)
	case "10":
		var fn = func(s, p string, as bool) {
			res := handle.IsMatch(s, p)
			fmt.Println(s, p, res == as)
		}
		fn("aa", "a", false)
		fn("aa", "a*", true)
		fn("ab", ".*", true)
		fn("aab", "c*a*b", true)
		fn("mississippi", "mis*is*p*.", false)
		// "mis*is*p*."
		//fn("abbbbbc", "a.*c", true)
		//fn("abbbbbc", "a.*c", true)
	case "12":

		fmt.Println(handle.IntToRoman(123) == "CXXIII")
		fmt.Println(handle.IntToRoman(3) == "III")
		fmt.Println(handle.IntToRoman(4) == "IV")
		fmt.Println(handle.IntToRoman(9) == "IX")
		fmt.Println(handle.IntToRoman(58) == "LVIII")
		fmt.Println(handle.IntToRoman(1994) == "MCMXCIV")
		fmt.Println(handle.IntToRoman(101) == "CI")
	case "13":
		fmt.Println(handle.RomanToInt("MCMXCIV"))
		fmt.Println(handle.RomanToInt("LVIII"))
	case "14":
		fmt.Println(handle.LongestCommonPrefix([]string{"aa", "aab", "aa"}))

		fmt.Println(handle.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
		fmt.Println(handle.LongestCommonPrefix([]string{"a"}))
	}

}
