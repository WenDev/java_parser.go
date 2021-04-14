package parser

/*
当前的解析步骤
根据状态转换图可以知道解析到某个步骤之后下一步需要解析什么,在此之前要先定义所有的步骤
*/
type state int

const (
	stateError                state = iota // 错误
	stateReservedOrIdentifier              // 保留字或者标识符
	stateNumber                            // 数字
	stateOperatorOrDelimiter               // 运算符
)
