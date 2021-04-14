package parser

import (
	"fmt"
	"regexp"
)

// 单词符号的类型,便于输出
type stringType int

const (
	RESERVED   = 1 // 保留字
	IDENTIFIER = 2 // 标识符
	NUMBER     = 3 // 数字
	OPERATOR   = 4 // 运算符
	DELIMITERS = 5 // 分隔符
	ERROR      = 6 // 错误字符
)

// isIdentifier 判断一个字符串是不是标识符
// @param s 要判断的字符串
// @return result 判断结果
func isIdentifier(s string) (result bool) {
	// 判断是否为保留字
	for _, word := range reservedWords {
		if s == word {
			return false
		}
	}

	// 标识符以下划线或数字开头，其余为字母、数字、下划线任意组合
	result, _ = regexp.MatchString("[a-zA-Z_][a-zA-Z_0-9]*", s)
	return
}

// printInfo 按格式要求输出读取到的单词符号的类型
// @param s 单词符号
// @param t 单词符号的类型
func printInfo(s string, t stringType) {
	fmt.Printf("(%d, \"%s\")\n", t, s)
}

// 返回两个数中较小的一个
func min(a, b int) (min int) {
	if a < b {
		return a
	} else {
		return b
	}
}
