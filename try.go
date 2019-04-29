package gotry

import (
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

func _Try(fn func()) (err *assert.KErr) {
	assert.Bool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	assert.Bool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			err = assert.NewKErr()
			if r := recover(); r != nil {
				switch d := r.(type) {
				case *assert.KErr:
					err = d
				case error:
					err.SetErr(d)
				}
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
