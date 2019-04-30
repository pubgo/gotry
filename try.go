package gotry

import (
	"errors"
	"github.com/pubgo/assert"
	"reflect"
)

type _try struct {
	err error
}

func (t *_try) Catch(fn func(err *assert.KErr)) {
	if t.err == nil {
		return
	}
	fn(t.err.(*assert.KErr))
}

func (t *_try) Error() error {
	return t.err
}

func (t *_try) P() {
	t.err.(*assert.KErr).LogStacks()
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
			m := &assert.KErr{}
			if r := recover(); r != nil {
				switch d := r.(type) {
				case *assert.KErr:
					m = d
				case error:
					m.Sub = d
				case string:
					m.Sub = errors.New(d)
				}
			}

			if m.Sub == nil {
				err = nil
			} else {
				err = m
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
