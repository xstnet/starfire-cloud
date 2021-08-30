package errors

import "errors"

func InvalidParameter() error {
	return New("参数错误")
}

func New(text string) error {
	return errors.New(text)
}
