package parser

/*
当前的解析步骤
根据状态转换图可以知道解析到某个步骤之后下一步需要解析什么,在此之前要先定义所有的步骤
*/
type state int

const (
	stateError state = iota // 错误,遇到错误就跳到这里并且停止解析
	stateA                  // 起始状态,状态A/B/D/E下读到保留字会进入此状态
	stateB                  // 状态A/B/D/E下读到标识符会进入此状态
	stateC                  // 状态A/D下读到常数会进入此状态
	stateD                  // 状态A/B/C/E下读到运算符会进入此状态
	stateE                  // 终结状态,状态A/B/C/D/E读到分隔符(界符)会进入此状态
)
