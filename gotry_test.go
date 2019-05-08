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
		assert.ErrWrap(ER, func(m *assert.M) {
			m.Msg("mmk")
			m.Tag("tag")
		})
	}).Catch(func(err error) {
		switch err {
		case ER:
			fmt.Println(err.Error())
			fmt.Println(err == ER)
		}
	}).CatchTag(func(tag string, err *assert.KErr) {
		fmt.Println(tag)
	})
}

type SS struct {
}

func (*SS) Error() string {
	return "ok" + time.Now().String()
}
func TestKind(t *testing.T) {
	gotry.Try(func() {
		assert.SWrap(&SS{}, "mmk")
		assert.ErrWrap(&SS{}, func(m *assert.M) {
			m.Msg("mmk")
		})

	}).Catch(func(err error) {
		switch err.(type) {
		case *SS:
		case error:
		}
		fmt.Println(err.Error())
	}).Finally(func(err *assert.KErr) {
		err.P()
	})
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
}
