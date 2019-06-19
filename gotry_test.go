package gotry_test

import (
	"errors"
	"fmt"
	"github.com/pubgo/assert"
	"github.com/pubgo/gotry"
	"testing"
	"time"
)

func TestNam12e(t *testing.T) {
	var ER = errors.New("dd")
	gotry.Try(func() {
		assert.SWrap(ER, func(m *assert.M) {
			m.Msg("mmk")
			m.Tag("tag")
		})
	}).Catch(func(err *assert.KErr) {
		switch err.Err() {
		case ER:
			fmt.Println(err.Error())
			fmt.Println(err == ER)
		}
	}).P()
}

type SS struct {
}

func (*SS) Error() string {
	return "ok" + time.Now().String()
}
func TestKind(t *testing.T) {
	gotry.Try(func() {
		assert.ErrWrap(&SS{}, "mmk")
		assert.SWrap(&SS{}, func(m *assert.M) {
			m.Msg("mmk")
		})

	}).Catch(func(err *assert.KErr) {
		switch err.Err().(type) {
		case *SS:
		case error:
		}
		fmt.Println(err.Error())
	}).P()
}

func hello(args ...string) bool {
	fmt.Println(args)

	for _, arg := range args {
		assert.TT(arg == "a", func(m *assert.M) {
			m.Msg("error panic  info")
			m.Tag("test")
		})
	}

	return true
}

func TestFn(t *testing.T) {
	gotry.Try(func() *SS {
		return &SS{}
	}).Then(func(vs *SS) string {
		return vs.Error()
	}).Then(func(s string) {
		fmt.Println(s)
	}).P()

	gotry.Try(fmt.Println, "test", 1, nil).
		Then(func(n int, err error) {
			fmt.Println(n, err)
		}).P()

	gotry.Try(hello, "ss", "ddd").Then(func(b bool) {
		fmt.Println(b)
	}).Catch(func(err *assert.KErr) {
		fmt.Println(err, "err")
	}).P()

	gotry.Try(hello, "ss", "ddd", "a").Then(func(b bool) {
		fmt.Println(b)
	}).Catch(func(err *assert.KErr) {
		fmt.Println(err, "err")
	}).P()

	gotry.Try(hello, "ss", "ddd", "a").Then(func(b bool) {
		fmt.Println(b)
	}).Expect("sss %s", "ss")

}

func Test_try_CatchTag(t *testing.T) {
	gotry.Try(hello, "ss", "ddd", "a").Then(func(b bool) {
		fmt.Println(b, "特斯特根哥哥哥哥")
	}).CatchTag("test", func(err *assert.KErr) {
		fmt.Println("test tag")
	}).P()
}
