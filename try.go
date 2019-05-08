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
	_t.err = _Try(func() {
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

func Try(f interface{}, params ...interface{}) *_try {
	t := reflect.TypeOf(f)
	assert.ST(t.Kind() != reflect.Func, "the params is not func type")

	_t := &_try{}
	_t.err = _Try(func() {
		var vs []reflect.Value
		for i, p := range params {
			if p == nil {
				if t.IsVariadic() {
					i = 0
				}

				vs = append(vs, reflect.New(t.In(i)).Elem())
			} else {
				vs = append(vs, reflect.ValueOf(p))
			}
		}
		_t._values = reflect.ValueOf(f).Call(vs)
	})
	return _t
}

func _Try(fn func()) (err error) {
	assert.ST(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	assert.ST(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			m := &assert.KErr{}
			if r := recover(); r != nil {
				switch d := r.(type) {
				case *assert.KErr:
					m = d
				case error:
					m.Err = d
					m.Msg = d.Error()
				default:
					panic("type error, must be *assert.KErr type")
				}
			}

			if m.Err == nil {
				err = nil
			} else {
				err = m
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
