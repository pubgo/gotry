package gotry

import (
	"time"
)

func Retry(num int, fn func() error) (err error) {
	for i := 0; i < num; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(time.Second * time.Second)
	}
	return
}

func Ticker(dur time.Duration, fn func(c int) bool) bool {
	for i := 0; fn(i); i++ {
		time.Sleep(dur)
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
