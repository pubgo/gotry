package gotry

import (
	"errors"
	"github.com/pubgo/assert"
	"reflect"
)

type _try struct {
	err    error
	params []reflect.Value
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

func Try(fn func()) *_try {
	assert.Bool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	assert.Bool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	t := &_try{}
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch d := r.(type) {
				case error:
					t.err = d
				case string:
					t.err = errors.New(d)
				}
			}
		}()
		t.params = reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return t
}
