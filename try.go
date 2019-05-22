package gotry

import (
	"reflect"
)

type _try struct {
	err     error
	_values []reflect.Value
}

func (t *_try) P() {
	_P(t.err)
}

func (t *_try) Panic() {
	if err := t.KErr(); err != nil {
		err.Caller = funcCaller()
		panic(err)
	}
}

func (t *_try) Then(fn interface{}) *_try {
	if t.err != nil && t.KErr().Err != nil {
		return t
	}

	_AssertFn(fn)

	_fn := reflect.ValueOf(fn)
	_ST(_fn.Type().NumIn() != len(t._values), "the params num is not match")

	_t := &_try{}
	_t.err = _KTry(func() {
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

func (t *_try) Expect(f string, args ...interface{}) {
	Try(_SWrap, t.err, func(m *_M) {
		m.Msg(f, args...)
		m.Tag("Expect")
	}).Finally(func(err *_KErr) {
		err.Caller = funcCaller()
		panic(err)
	});
}

// real error
func (t *_try) CatchTag(tag string, fn func(err *_KErr)) *_try {
	_ke := t.KErr()
	if t.err == nil || len(t._values) != 0 || _ke.Tag == "" {
		return t
	}

	if _err := t.Err(); _err == nil {
		return t
	}

	if _ke.Tag == tag {
		fn(_ke)
	}

	return t
}

func (t *_try) Finally(fn func(err *_KErr)) {
	if t.err == nil {
		return
	}

	fn(t.err.(*_KErr))
}

// real err
func (t *_try) Err() error {
	if err := t.KErr(); err != nil {
		return err.Err
	}
	return nil
}

// wrap err
func (t *_try) KErr() *_KErr {
	if t.err == nil {
		return nil
	}
	return t.err.(*_KErr)
}

func Try(f interface{}, args ...interface{}) *_try {
	_t := &_try{}
	_t.err = _KTry(func() {
		_t._values = _FnOf(f, args...)()
	})
	return _t
}
