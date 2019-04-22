package gotry_test

import (
	"fmt"
	"github.com/pubgo/assert"
	"github.com/pubgo/gotry"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	tsk := gotry.NewTask(1000000, time.Second, func(i int) {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 999)
	})

	tt := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		tsk.Do(i)
	}

	fmt.Println(time.Now().UnixNano() - tt)
}

func TestWaitFor(t *testing.T) {
	assert.P(gotry.WaitFor(func(c time.Duration) bool {
		assert.Bool(c > time.Second*time.Duration(10), "")
		return true
	}))
}
