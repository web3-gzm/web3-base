package base

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：

	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	每个右括号都有一个对应的相同类型的左括号。

	示例 1：
	输入：s = "()"
	输出：true

	示例 2：
	输入：s = "()[]{}"
	输出：true

	示例 3：
	输入：s = "(]"
	输出：false

	示例 4：
	输入：s = "([])"
	输出：true

	示例 5：
	输入：s = "([)]"
	输出：false
*/
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	temp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{'}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if data, ok := temp[s[i]]; ok && data == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}

	return len(stack) == 0
}

/*
*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = getMaxPrefix(strs[i], strs[i-1])
		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}
func getMaxPrefix(s1, s2 string) string {
	minLen := min(len(s1), len(s2))
	resultLeg := 0
	for resultLeg < minLen && s1[resultLeg] == s2[resultLeg] {
		resultLeg++
	}

	return s1[:resultLeg]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
