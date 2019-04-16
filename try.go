package try

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

func Try(fn func(assert *assert.Assert)) *_try {
	assert.Bool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	assert.Bool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	t := &_try{}
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch r.(type) {
				case error:
					t.err = r.(error)
				case string:
					t.err = errors.New(r.(string))
				}
			}
		}()
		t.params = reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(_assert)})
	}()
	return t
}
