package gotry

import (
	"github.com/pubgo/assert"
	"log"
	"reflect"
	"runtime"
	"time"
)

func NewTask(max int, maxDur time.Duration) *task {
	_t := &task{max: max, maxDur: maxDur, q: make(chan *_task_fn, max)}
	_t._handle()
	return _t
}

type _task_fn struct {
	fn   interface{}
	args []interface{}
}

type task struct {
	maxDur time.Duration
	curDur time.Duration
	max    int
	q      chan *_task_fn
}

func (t *task) Do(f interface{}, args ...interface{}) {
	assert.Bool(f == nil || reflect.TypeOf(f).Kind() != reflect.Func, "please init params")

	for {
		if len(t.q) < t.max && t.curDur < t.maxDur {
			t.q <- &_task_fn{
				fn:   f,
				args: args,
			}
			break
		}

		if len(t.q) < runtime.NumCPU()*2 {
			t.curDur = 0
		}

		log.Printf("q_l:%d cur_dur:%s", len(t.q), t.curDur.String())
		time.Sleep(time.Millisecond * 200)
	}

}

func FnCost(f func()) time.Duration {
	t1 := time.Now()
	f()
	return time.Now().Sub(t1)
}

func (t *task) _handle() {
	go func() {
		for {
			select {
			case _fn := <-t.q:
				go func() {
					t.curDur = FnCost(func() {
						Fn(_fn.fn, _fn.args...).Assert()
					})
				}()
			case <-time.NewTicker(time.Second * 2).C:

			}
		}
	}()
}
