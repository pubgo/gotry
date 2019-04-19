package gotry_test

import (
	"fmt"
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
