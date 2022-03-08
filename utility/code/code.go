package code

import "fmt"

type localCode struct {
	code    int    // 错误码
	message string // 提示信息
}

var (
	Ok  = localCode{0, "成功"}
	Nil = localCode{-1, "失败"}
)

// 返回错误码
func (c localCode) Code() int {
	return c.code
}

// 返回错误码的提示信息
func (c localCode) Message() string {
	return c.message
}

// 以字符串的形式返回错误码
func (c localCode) String() string {
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
