package gotry

import (
	"github.com/pubgo/assert"
	"reflect"
)

type _try struct {
	err     error
	_values []reflect.Value
}

func (t *_try) P() {
	assert.P(t.err)
}

func (t *_try) Then(fn interface{}) *_try {
	if t.err != nil || len(t._values) == 0 {
		return t
	}

	_fn := reflect.ValueOf(fn)
	assert.ST(_fn.Kind() != reflect.Func, "the params is not func type")
	assert.ST(_fn.Type().NumIn() != len(t._values), "the params num is not match")

	_t := &_try{}
	_t.err = assert.KTry(func() {
		_t._values = _fn.Call(t._values)
	})

	return _t
}

// real error
func (t *_try) Catch(fn func(err error)) *_try {
	if t.err == nil || len(t._values) != 0 {
		return t
	}

	_err := t.Err()
	if _err == nil {
		return t
	}

	fn(_err)
	return t
}

// real error
func (t *_try) CatchTag(fn func(tag string, err *assert.KErr)) *_try {
	if t.err == nil || len(t._values) != 0 || t.KErr().Tag == "" {
		return t
	}

	_err := t.Err()
	if _err == nil {
		return t
	}

	fn(t.KErr().Tag, t.KErr())
	return t
}

func (t *_try) Finally(fn func(err *assert.KErr)) {
	if t.err == nil {
		return
	}

	fn(t.KErr())
}

func (t *_try) Err() error {
	if err := t.KErr(); err != nil {
		return err.Err
	}
	return nil
}

func (t *_try) KErr() *assert.KErr {
	if t.err == nil {
		return nil
	}
	return t.err.(*assert.KErr)
}

func Try(f interface{}, args ...interface{}) *_try {
	_t := &_try{}
	_t.err = assert.KTry(func() {
		_t._values = assert.FnOf(f, args...)()
	})
	return _t
}
