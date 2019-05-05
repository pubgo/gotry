package gotry_test

import (
	"errors"
	"fmt"
	"github.com/pubgo/assert"
	"github.com/pubgo/gotry"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	tsk := gotry.NewTask(1000000, time.Second*2)

	tt := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		//tsk.Do(i)
		tsk.Do(i)
	}

	fmt.Println(time.Now().UnixNano() - tt)
}

func TestWaitFor(t *testing.T) {
	gotry.Try(func() {
		assert.Err(gotry.WaitFor(func(c time.Duration) bool {
			fmt.Println(c)
			assert.Bool(c > time.Second*time.Duration(10), "")
			return true
		}))
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	})

}

func TestClock(t *testing.T) {
	fmt.Println(time.Now().Clock())
	gotry.Try(func() {
		assert.Err(gotry.Ticker(func(dur time.Time) time.Duration {
			fmt.Println(dur.Clock())
			return time.Second * 10
		}))
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	})
}

func TestNam12e(t *testing.T) {
	gotry.Try(func() {
		assert.ErrWrap(errors.New("dd"), "mmk")
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	})
}

type SS struct {
}

func (*SS) Error() string {
	return "ok"
}
func TestKind(t *testing.T) {
	gotry.Try(func() {
		//assert.ErrWrap(errors.New("sss"), "mmk")
		assert.ErrWrap(&SS{}, "mmk")
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
