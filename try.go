package gotry

import (
	"errors"
	"fmt"
	"github.com/pubgo/assert"
	"reflect"
)

type _try struct {
	err error
}

func (t *_try) Catch(fn func(err error)) {
	if t.err == nil {
		return
	}
	fn(t.err)
}

func (t *_try) Error() error {
	return t.err
}

func (t *_try) P() {
	if t.err != nil {
		fmt.Println(t.err)
	}
}

func Try(fn func()) *_try {
	return &_try{err: _Try(fn)}
}

func _Try(fn func()) (err error) {
	assert.Bool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	assert.Bool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch d := r.(type) {
				case error:
					err = d
				case string:
					err = errors.New(d)
				}
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
