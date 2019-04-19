package gotry

import (
	"github.com/pubgo/assert"
	"time"
)

func Retry(num int, fn func() error) (err error) {
	for i := 0; i < num; i++ {
		if err = Try(func() {
			assert.MustNotError(fn())
		}).Error(); err == nil {
			return nil
		}
		time.Sleep(time.Second * time.Second)
	}
	return
}

func Ticker(dur time.Duration, fn func(dur time.Duration) bool) bool {
	_dur := time.Duration(0)
	for fn(_dur) {
		time.Sleep(dur)
		_dur += dur
	}
	return false
}

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}
