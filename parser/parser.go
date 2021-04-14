package parser

import (
	"regexp"
)

type parser struct {
	java     string // 待解析的Java语句
	position int    // 当前解析到的位置
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

	// 判断保留字
	for _, rw := range reservedWords {
		token := p.java[p.position:min(len(p.java), p.position+len(rw))]
		if token == rw {
			return token, len(token)
		}
	}

	// 不是保留字,peek其他子句
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


// 弹出所有空格
func (p *parser) popWhitespace() {
	for ; p.position < len(p.java) && p.java[p.position] == ' '; p.position++ {
	}
}
