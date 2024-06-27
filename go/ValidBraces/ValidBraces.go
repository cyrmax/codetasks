/*
Valid braces
From codewars.com

Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.
This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!
All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

What is considered Valid?
A string of braces is considered valid if all braces are matched with the correct brace.

Examples
"(){}[]"   =>  True
"([{}])"   =>  True
"(}"       =>  False
"[(])"     =>  False
"[({})](]" =>  False
*/

package kata

func ValidBraces(str string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(', '[', '{':
			stack = append(stack, string(str[i]))
		case ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
