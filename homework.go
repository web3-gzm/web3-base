package homeword01

import (
	"fmt"
	"sort"
)

func main() {
	i := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(i)
}

/*
1. 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*/
func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	temp := map[int]int{}

	for _, v := range nums {
		temp[v] = temp[v] + 1
	}

	for k, v := range temp {
		if v == 1 {
			return k
		}
	}

	return 0
}

/*
2. 回文数
判断一个整数是否是回文数
121
*/
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	if x < 10 {
		return true
	}

	temp := 0

	for x > temp {
		last := x % 10
		temp = temp*10 + last
		x = x / 10
	}

	return x == temp || x == temp/10
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	if s == "" || len(s)%2 != 0 {
		return false
	}
	// 维护一个map，存储左括号和右括号
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			if len(stack) == 0 {
				return false
			}
			// 如果找到右括号
			last := stack[len(stack)-1]
			if last != v {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, s[i])
	}

	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = getMinPrefix(prefix, strs[i])
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
func getMinPrefix(str1, str2 string) string {
	minLength := getMinLength(str1, str2)

	maxLengPrefix := 0
	for minLength > maxLengPrefix && str1[maxLengPrefix] == str2[maxLengPrefix] {
		maxLengPrefix++
	}

	return str1[:maxLengPrefix]
}

func getMinLength(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return 0
	}
	if len(str1) > len(str2) {
		return len(str2)
	}

	return len(str1)
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

func PlusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		temp := digits[i] + carry
		digits[i] = temp % 10
		carry = temp / 10
		if carry == 0 {
			break
		}
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		// 当前区间的尾部小于下一个区间的头部,说明不重叠，直接加入即可
		if last[1] < intervals[i][0] {
			ans = append(ans, intervals[i])
		} else {
			// 需要合并区间
			last[1] = max(last[1], intervals[i][1])
		}
	}

	return ans
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	result := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if v, ok := result[target-nums[i]]; ok {
			return []int{v, i}
		}
		result[nums[i]] = i
	}
	return nil
}
