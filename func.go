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

func Retry(num int, fn func() error) error {
	_t := fibonacci()
	for i := 0; i < num; i++ {
		if err := _Try(func() {
			assert.Err(fn())
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * time.Duration(_t()))
	}
	return nil
}

func WaitFor(fn func(dur time.Duration) bool) error {
	var _b = true
	for i := 0; _b; i++ {
		if err := _Try(func() {
			_b = fn(time.Second * time.Duration(i))
		}); err != nil {
			return err
		}

		if !_b {
			return nil
		}

		time.Sleep(time.Second)
	}
	return nil
}

func Ticker(fn func(dur time.Time) time.Duration) error {
	_dur := time.Duration(0)
	for i := 0; ; i++ {
		if err := _Try(func() {
			_dur = fn(time.Now())
		}); err != nil {
			return err
		}

		if _dur < 0 {
			return nil
		}

		if _dur == 0 {
			_dur = time.Second
		}

		time.Sleep(_dur)
	}
}

func FnCost(f func()) time.Duration {
	t1 := time.Now()
	f()
	return time.Now().Sub(t1)
}