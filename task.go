package gotry

import (
	"log"
	"runtime"
	"time"
)

func NewTask(max int, maxDur time.Duration, fn interface{}) *task {
	return &task{max: max, maxDur: maxDur, q: make(chan int, max), handle: fn}
}

type task struct {
	maxDur time.Duration
	curDur time.Duration
	max    int
	q      chan int
	handle interface{}
}

func (t *task) Do(i ...interface{}) {
	for {
		if len(t.q) < t.max && t.curDur < t.maxDur {
			break
		}

		if len(t.q) < runtime.NumCPU()*2 {
			t.curDur = 0
		}

		log.Printf("q_l:%d cur_dur:%s", len(t.q), t.curDur)
		time.Sleep(time.Millisecond * 200)
	}

	go func() {
		t1 := time.Now()
		t.q <- 1
		Fn(t.handle, i...).Assert()
		<-t.q
		t.curDur = time.Now().Sub(t1)
	}()
}
