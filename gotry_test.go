package gotry_test

import (
	"errors"
	"fmt"
	"github.com/pubgo/assert"
	"github.com/pubgo/gotry"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	gotry.Try(func() {
		assert.Throw(gotry.Retry(3, func() {
			fmt.Println("ok")
			assert.T(true, func(m *assert.M) {
				m.Msg("test retry")
			})
		}))

	}).Catch(func(err error) {
		fmt.Println(err.Error())
	}).P()

	fmt.Println(gotry.Try(func() {
		assert.Throw(gotry.Retry(3, func() {
			fmt.Println("ok")

			assert.T(true, func(m *assert.M) {
				m.Msg("test retry")
			})
		}))
	}).Err())
}

func TestTask(t *testing.T) {
	var ss = gotry.FuncOf(func(i int) {
		assert.T(i > 10000, func(m *assert.M) {
			m.Msg("max index")
		})

		fmt.Println(i)
	}, func(err error) {
		fmt.Println(err)
		fmt.Println(err.(*assert.KErr).Err.Error() == errors.New("max index").Error())
	})

	tsk := gotry.NewTask(1000000, time.Second*2)

	tt := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		tsk.Do(ss, i)
	}

	fmt.Println(time.Now().UnixNano() - tt)
}

func TestWaitFor(t *testing.T) {
	gotry.Try(func() {
		gotry.WaitFor(func(c time.Duration) bool {
			fmt.Println(c)
			assert.T(c > time.Second*time.Duration(2), func(m *assert.M) {
				m.Msg("time out")
			})

			return true
		})
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	}).P()

}

func TestClock(t *testing.T) {
	fmt.Println(time.Now().Clock())
	gotry.Try(func() {
		gotry.Ticker(func(dur time.Time) time.Duration {
			fmt.Println(dur.Clock())
			assert.T(true, func(m *assert.M) {
				m.Msg("ddd test")
			})
			return time.Second * 1
		})
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	}).P()
}

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
