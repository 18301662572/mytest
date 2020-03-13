package convey

import (
	"errors"
)

//convey 单元测试框架

func Func(arg string) error {
	if len(arg) > 0 {
		return nil
	} else {
		return errors.New("arg is nil")
	}
}


