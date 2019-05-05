package gotry

import (
	"github.com/pubgo/assert"
	"reflect"
)

type _try struct {
	err error
}

// real error
func (t *_try) Catch(fn func(err error)) *_try {
	if t.err == nil {
		return t
	}

	_err := t.Error()
	if _err == nil {
		return t
	}

	fn(_err)
	return t
}

func (t *_try) Finally(fn func(err *assert.KErr)) {
	if t.err == nil {
		return
	}

	fn(t.KError())
}

func (t *_try) Error() error {
	if err := t.KError(); err != nil {
		return err.Err
	}

	return nil
}

func (t *_try) KError() *assert.KErr {
	if t.err == nil {
		return nil
	}

	return t.err.(*assert.KErr)
}

func (t *_try) P() {
	assert.P(t.err)
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
				default:
					panic("type error, must be *assert.KErr type")
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
