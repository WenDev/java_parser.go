package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type parser struct {
	java     string // 待解析的Java语句
	position int    // 当前解析到的位置
	err      error  // 解析过程中出现的错误
	state    state  // 当前状态
}

// 弹出解析的下一个记号
func (p *parser) pop() (peeked string) {
	// 得到下一个记号
	peeked, length := p.peekWithLength()
	// 把当前解析位置移动到下一个记号后
	p.position += length
	// 把空格全部弹出
	p.popWhitespace()

	return peeked
}

// 返回但不弹出解析的下一个记号
func (p *parser) peek() (peeked string) {
	// 返回下一个记号（这里不需要长度，pop才需要）
	peeked, _ = p.peekWithLength()
	return peeked
}

// 返回读到的字句及其长度
func (p *parser) peekWithLength() (string, int) {
	// 读到末尾,返回空串
	if p.position >= len(p.java) {
		return "", 0
	}

	// 判断界符
	for _, d := range delimiters {
		token := p.java[p.position:p.position+1]
		if d == token {
			return token, 1
		}
	}

	// 判断操作符
	for _, o := range operators {
		token := p.java[p.position:p.position+2]
		if strings.Contains(token, o) {
			return o, len(o)
		}
	}

	// 判断保留字
	for _, rw := range reservedWords {
		token := p.java[p.position:min(len(p.java), p.position+len(rw))]
		if token == rw {
			return token, len(token)
		}
	}

	// 判断数字
	if n, l := p.peekNumber(); l != 0 {
		return n, l
	}

	// 不是保留字也不是数字,peek其他子句
	return p.peekIdentifierWithLength()
}

func (p *parser) peekIdentifierWithLength() (identifier string, length int) {
	// 不在语句的最后,根据标识符的构成规则进行匹配,返回对应的单词符号及其长度
	for i := p.position; i < len(p.java); i++ {
		if matched, _ := regexp.MatchString(`[a-zA-Z0-9_*]`, string(p.java[i])); !matched {
			return p.java[p.position:i], len(p.java[p.position:i])
		}
	}

	// 在语句的最后,截取当前位置到最后的部分,返回被截取的部分及其长度
	return p.java[p.position:], len(p.java[p.position:])
}

func (p *parser) peekNumber() (number string, length int) {
	for i := p.position; i < len(p.java); i++ {
		if matched, _ := regexp.MatchString(`^[+-]?([0-9]*\.?[0-9]+|[0-9]+\.?[0-9]*)([eE][+-]?[0-9]+)?$`, string(p.java[i])); !matched {
			return p.java[p.position:i], len(p.java[p.position:i])
		}
	}

	return "", 0
}

// 弹出所有空格
func (p *parser) popWhitespace() {
	for ; p.position < len(p.java) && (p.java[p.position] == ' ' || p.java[p.position] == '\n'); p.position++ {
	}
}

// Parse 提供给外部进行调用的函数
func Parse(java string) (err error) {
	err = parse(java)
	return err
}

// parse 构造一个自动机Parser,开始解析
func parse(java string) (err error) {
	return (&parser{
		java:     strings.TrimSpace(java),
		position: 0,
		err:      nil,
		state:    stateA, // 状态A为起始状态
	}).parse()
}

// parse 解析的入口函数
func (p *parser) parse() (err error) {
	err = p.doParse()
	return err
}

// doParse 主解析函数
func (p *parser) doParse() (err error) {
	for {
		if p.position >= len(p.java) {
			return p.err
		}
		switch p.state {
		case stateA:
			reserved := p.peek()
			if isReserved(reserved) {
				printInfo(reserved, RESERVED)
				p.pop()
			} else {
				// 不是标识符,出错
				p.err = fmt.Errorf("Except a reserve word, got '%s' ", reserved)
				p.state = stateError
				continue
			}

			nextToken := p.peek()
			if isReserved(nextToken) {
				p.state = stateA
			} else if isIdentifier(nextToken) {
				p.state = stateB
			} else if isOperator(nextToken) {
				p.state = stateD
			} else if isDelimiter(nextToken) {
				p.state = stateE
			} else if isNumber(nextToken) {
				p.state = stateC
			}
		case stateB:
			identifier := p.peek()
			if isIdentifier(identifier) {
				printInfo(identifier, IDENTIFIER)
				p.pop()
			} else {
				p.err = fmt.Errorf("Except an identifier, got '%s' ", identifier)
				p.state = stateError
			}

			nextToken := p.peek()
			if isIdentifier(nextToken) {
				p.state = stateB
			} else if isOperator(nextToken) {
				p.state = stateD
			} else if isDelimiter(nextToken) {
				p.state = stateE
			} else if isReserved(nextToken) {
				p.state = stateA
			}
		case stateC:
			number := p.peek()
			if isNumber(number) {
				printInfo(number, NUMBER)
				p.pop()
			} else {
				p.err = fmt.Errorf("Except a number, got '%s' ", number)
			}

			nextToken := p.peek()
			if isDelimiter(nextToken) {
				p.state = stateE
			} else if isOperator(nextToken) {
				p.state = stateD
			}
		case stateD:
			operator := p.peek()
			if isOperator(operator) {
				printInfo(operator, OPERATOR)
				p.pop()
			} else {
				p.err = fmt.Errorf("Except an operator, got '%s' ", operator)
			}

			nextToken := p.peek()
			if isNumber(nextToken) {
				p.state = stateC
			} else if isDelimiter(nextToken) {
				p.state = stateE
			} else if isIdentifier(nextToken) {
				p.state = stateB
			} else if isReserved(nextToken) {
				p.state = stateA
			}
		case stateE:
			delimiter := p.peek()
			if isDelimiter(delimiter) {
				printInfo(delimiter, DELIMITERS)
				p.pop()
			} else {
				p.err = fmt.Errorf("Except a delimiter, got '%s' ", delimiter)
			}

			// 终结状态
			if p.position >= len(p.java) {
				return p.err
			}

			nextToken := p.peek()
			if isReserved(nextToken) {
				p.state = stateA
			} else if isIdentifier(nextToken) {
				p.state = stateB
			} else if isNumber(nextToken) {
				p.state = stateC
			} else if isDelimiter(nextToken) {
				p.state = stateE
			} else if isOperator(nextToken) {
				p.state = stateD
			}
		case stateError:
			return p.err
		default:
			return nil
		}
	}
}
