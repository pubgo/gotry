package gotry

import (
	"errors"
	"fmt"
	"github.com/pubgo/assert"
	"testing"
	"time"
)

func TestNam12e(t *testing.T) {
	var ER = errors.New("dd")
	Try(func() {
		_SWrap(ER, func(m *_M) {
			m.Msg("mmk")
			m.Tag("tag")
		})
	}).Catch(func(err error) {
		switch err {
		case ER:
			fmt.Println(err.Error())
			fmt.Println(err == ER)
		}
	})
}

type SS struct {
}

func (*SS) Error() string {
	return "ok" + time.Now().String()
}
func TestKind(t *testing.T) {
	Try(func() {
		_ErrWrap(&SS{}, "mmk")
		_SWrap(&SS{}, func(m *_M) {
			m.Msg("mmk")
		})

	}).Catch(func(err error) {
		switch err.(type) {
		case *SS:
		case error:
		}
		fmt.Println(err.Error())
	}).Finally(func(err *_KErr) {
		err.P()
	})
}

func hello(args ...string) bool {
	fmt.Println(args)

	for _, arg := range args {
		_TT(arg == "a", func(m *assert.M) {
			m.Msg("error panic  info")
			m.Tag("test")
		})
	}

	return true
}

func TestFn(t *testing.T) {
	Try(func() *SS {
		return &SS{}
	}).Then(func(vs *SS) string {
		return vs.Error()
	}).Then(func(s string) {
		fmt.Println(s)
	}).P()

	Try(fmt.Println, "test", 1, nil).
		Then(func(n int, err error) {
			fmt.Println(n, err)
		}).P()

	Try(hello, "ss", "ddd").Then(func(b bool) {
		fmt.Println(b)
	}).Catch(func(err error) {
		fmt.Println(err, "err")
	}).Finally(func(err *_KErr) {
		err.P()
	})

	Try(hello, "ss", "ddd", "a").Then(func(b bool) {
		fmt.Println(b)
	}).Catch(func(err error) {
		fmt.Println(err, "err")
	}).Finally(func(err *_KErr) {
		err.P()
	})

	Try(func() {
		Try(hello, "ss", "ddd", "a").Then(func(b bool) {
			fmt.Println(b)
		}).Expect("sss %s", "ss")
	}).P()

}

func Test_try_CatchTag(t *testing.T) {
	Try(hello, "ss", "ddd", "a").Then(func(b bool) {
		fmt.Println(b,"特斯特根哥哥哥哥")
	}).CatchTag("test", func(err *_KErr) {
		fmt.Println("test tag",err.StackTrace())
	}).P()
}
