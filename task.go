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

func FuncOf(fn interface{}, efn func(err error)) TaskFn {
	t := reflect.TypeOf(fn)
	assert.ST(t.Kind() != reflect.Func, "the params is not func type")

	return func(args ...interface{}) *_task_fn {
		return &_task_fn{
			fn:   fn,
			args: args,
			efn:  efn,
		}
	}
}

type _task_fn struct {
	fn   interface{}
	args []interface{}
	efn  func(err error)
}

func (t *_task_fn) _do() {
	Try(t.fn, t.args...).Finally(func(err *assert.KErr) {
		if t.efn != nil {
			t.efn(err)
		}
	})
}

type task struct {
	maxDur time.Duration
	curDur time.Duration
	max    int
	q      chan *_task_fn
}

func (t *task) Do(f TaskFn, args ...interface{}) {

	for {
		if len(t.q) < t.max && t.curDur < t.maxDur {
			t.q <- f(args...)
			break
		}

		if len(t.q) < runtime.NumCPU()*2 {
			t.curDur = 0
		}

		log.Printf("q_l:%d cur_dur:%s", len(t.q), t.curDur.String())
		time.Sleep(time.Millisecond * 200)
	}
}

func (t *task) _handle() {
	go func() {
		for {
			select {
			case _fn := <-t.q:
				go func() {
					t.curDur = FnCost(_fn._do)
				}()
			case <-time.NewTicker(time.Second * 2).C:

			}
		}
	}()
}
