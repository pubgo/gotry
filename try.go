package gotry

import (
	"github.com/pubgo/assert"
	"os"
	"reflect"
)

type _try struct {
	err     error
	_values []reflect.Value
}

func (t *_try) P() {
	if err := t.KErr(); err != nil {
		err.P()
	}
}

func (t *_try) Panic() {
	if err := t.KErr(); err != nil {
		err.Caller(funcCaller())
		panic(err)
	}
}

func (t *_try) Then(fn interface{}) *_try {
	if t.err != nil && t.KErr().Err() != nil {
		return t
	}

	assert.AssertFn(fn)

	_fn := reflect.ValueOf(fn)
	assert.T(_fn.Type().NumIn() != len(t._values), "the params num is not match")

	_t := &_try{}
	_t.err = assert.KTry(func() {
		_t._values = _fn.Call(t._values)
	})

	return _t
}

// real error
func (t *_try) Catch(fn func(err *assert.KErr)) *_try {
	if t.err == nil || len(t._values) != 0 {
		return t
	}

	fn(t.KErr())
	return t
}

func (t *_try) Expect(f string, args ...interface{}) {
	Try(assert.SWrap, t.err, func(m *assert.M) {
		m.Msg(f, args...)
		m.Tag("Expect")
	}).Catch(func(err *assert.KErr) {
		err.Caller(funcCaller())
		err.P()
		os.Exit(1)
	})
}

// tag error
func (t *_try) CatchTag(tag string, fn func(err *assert.KErr)) *_try {
	_ke := t.KErr()
	if t.err == nil || len(t._values) != 0 || _ke.Tag() == "" {
		return t
	}

	if _err := t.Err(); _err == nil {
		return t
	}

	if _ke.Tag() == tag {
		fn(_ke)
	}

	return t
}

// real err
func (t *_try) Err() error {
	if err := t.KErr(); err != nil {
		return err.Err()
	}
	return nil
}

// wrap err
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
