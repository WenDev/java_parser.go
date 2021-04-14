package parser

// reservedWords Java中的保留字
var reservedWords = []string {
	"abstract",
	"assert",
	"boolean",
	"break",
	"byte",
    "case",
	"catch",
	"char",
	"class",
	"const",
	"continue",
	"default",
	"do",
	"double",
	"else",
	"enum",
	"extends",
	"final",
	"finally",
	"float",
	"for",
	"goto",
	"if",
	"implements",
	"import",
	"instanceof",
	"int",
	"interface",
	"long",
	"native",
	"new",
	"package",
	"private",
	"protected",
	"public",
	"return",
	"strictfp",
	"short",
	"static",
	"super",
	"switch",
	"synchronized",
	"this",
	"throw",
	"throws",
	"transient",
	"try",
	"void",
	"volatile",
	"while",
}

// delimiters 界符
var delimiters = []string{
	"(",
	")",
	"[",
	"]",
	"{",
	"}",
	"\"",
	"'",
	"<",
	">",
}

// operators 运算符
var operators = []string {
	// ========== 算术运算符 ==========
	"+",
	"-",
	"*",
	"/",
	"%",
	"++",
	"--",
	// ========== 关系运算符 ==========
	"==",
	"!=",
	">",
	">=",
	"<",
	"<=",
	// ========== 赋值运算符 ==========
	"=",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"<<=",
	">>=",
	"&=",
	"^=",
	"|=",
	// ========== 位运算符 ==========
	"&",
	"|",
	"^",
	"~",
	"<<",
	">>",
	">>>",
	// ========== 逻辑运算符 ==========
	"&&",
	"||",
	"!",
	// ========== 点运算符 ==========
	".",
}