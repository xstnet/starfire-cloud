package errors

func SystemError() error {
	return New("系统错误")
}
