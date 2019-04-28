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
	tsk := gotry.NewTask(1000000, time.Second*2, func(i int) {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 999)
	})

	tt := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		//tsk.Do(i)
		tsk.GoDo(i)
	}

	fmt.Println(time.Now().UnixNano() - tt)
}

func TestWaitFor(t *testing.T) {
	gotry.Try(func() {
		assert.MustNotError(gotry.WaitFor(func(c time.Duration) bool {
			fmt.Println(c)
			assert.Bool(c > time.Second*time.Duration(10), "")
			return true
		}))
	}).Catch(func(err *assert.KErr) {
		fmt.Println(err.Error())
	})

}

func TestClock(t *testing.T) {
	fmt.Println(time.Now().Clock())
	gotry.Try(func() {
		assert.MustNotError(gotry.Ticker(func(dur time.Time) time.Duration {
			fmt.Println(dur.Clock())
			return time.Second * 10
		}))
	}).Catch(func(err *assert.KErr) {
		fmt.Println(err.Error())
		err.LogStacks()
	})
}

func TestNam12e(t *testing.T) {
	gotry.Try(func() {
		assert.Err(errors.New("dd"), "mmk")
	}).Catch(func(err *assert.KErr) {
		fmt.Println(err.Error())
		fmt.Println(err.GetStacks())
	})
}
