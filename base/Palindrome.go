package base

/*
*
输入：nums = [2,2,1]给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。

示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
*/

/*
*
解题思路：反转数据

	从x尾部每次截取一位，放到一个新的变量中，同时截取x最后一位，直到新的变量的数字大于x，
	此时就将数据的一半进行翻转了，如果是奇数则需要将新变量左移一位，如果是偶数则直接判断是否相等
*/
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	result := 0
	for x > result {
		result = result*10 + x%10
		x /= 10
	}
	return x == result || x == result/10
}

/*
*
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
