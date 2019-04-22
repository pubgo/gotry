package gotry

import (
	"github.com/pubgo/assert"
	"time"
)

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}

func Retry(num int, fn func() error) {
	_t := fibonacci()
	for i := 0; i < num; i++ {
		if err := _Try(func() {
			assert.MustNotError(fn())
		}); err == nil {
			return
		}
		time.Sleep(time.Second * time.Duration(_t()))
	}
	return
}

func WaitFor(fn func(dur time.Duration) bool) (b bool) {
	var _b = true
	for i := 0; _b; i++ {
		if err := _Try(func() {
			_b = fn(time.Second * time.Duration(i))
		}); err != nil {
			return false
		}

		if !_b {
			return false
		}

		time.Sleep(time.Second)
	}
	return false
}

func Ticker(fn func(dur time.Time) time.Duration) {
	_dur := time.Duration(0)
	for i := 0; ; i++ {
		if err := _Try(func() {
			_dur = fn(time.Now())
		}); err != nil {
			return
		}

		if _dur < 0 {
			return
		}

		if _dur == 0 {
			_dur = time.Second
		}

		time.Sleep(_dur)
	}
}
