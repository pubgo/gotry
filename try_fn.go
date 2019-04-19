package gotry

import (
	"github.com/pubgo/assert"
	"reflect"
)

type _try_fn struct {
	err     error
	_values []reflect.Value
}

func (t *_try_fn) Assert() {
	assert.MustNotError(t.err)
}

func (t *_try_fn) Error() error {
	return t.err
}

func (t *_try_fn) Then(fn func(vs ...interface{})) *_try {
	if t.err != nil {
		return &_try{err: t.err}
	}

	_fn := reflect.ValueOf(fn)
	assert.Bool(_fn.Kind() != reflect.Func, "err -> Wrap: please input func")

	return Try(func() {
		_fn.Call(t._values)
	})
}

func Fn(f interface{}, params ...interface{}) *_try_fn {
	t := reflect.TypeOf(f)
	assert.Bool(t.Kind() != reflect.Func, "err -> Wrap: please input func")

	_t := &_try_fn{}
	_t.err = _Try(func() {
		var vs []reflect.Value
		for i, p := range params {
			if p == nil {
				vs = append(vs, reflect.New(t.In(i)).Elem())
			} else {
				vs = append(vs, reflect.ValueOf(p))
			}
		}
		_t._values = reflect.ValueOf(f).Call(vs)
	})
	return _t
}
