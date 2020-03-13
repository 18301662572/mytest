package convey

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//convey 单元测试框架

func TestFunc(t *testing.T) {
	Convey("test Func", t, func() {
		Convey("Func should return nil when arg is not empty", func() {
			arg := "1"
			err := Func(arg)
			So(err, ShouldBeNil)
		})
		Convey("Func should return error when arg is empty", func() {
			arg := ""
			exceptErr := errors.New("arg is nil")
			err := Func(arg)
			So(err, ShouldBeError, exceptErr)
		})
	})
}
