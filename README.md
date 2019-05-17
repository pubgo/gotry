# gotry
gotry for go

```go
var ER = errors.New("dd")
Try(func() {
  _ErrWrap(ER, func(m *_M) {
    m.Msg("mmk")
    m.Tag("tag")
  })
}).Catch(func(err error) {
  switch err {
  case ER:
    fmt.Println(err.Error())
    fmt.Println(err == ER)
  }
}).CatchTag(func(tag string, err *_KErr) {
  fmt.Println(tag)
})
```

```go
Try(func() {
  _SWrap(&SS{}, "mmk")
  _ErrWrap(&SS{}, func(m *_M) {
        m.Msg("mmk")
  })

}).Catch(func(err error) {
  switch err.(type) {
  case *SS:
  case error:
  }
fmt.Println(err.Error())
}).Finally(func(err *_KErr) {
	err.P()
})
```

```go
Try(func() *SS {
	return &SS{}
}).Then(func(vs *SS) string {
	return vs.Error()
}).Then(func(s string) {
	fmt.Println(s)
}).P()

Try(fmt.Println, "test", 1, nil).
	Then(func(n int, err error) {
	fmt.Println(n, err)
}).P()
```


```go
func hello(args ...string) bool {
	fmt.Println(args)

	for _, arg := range args {
		if arg == "a" {
			panic("error panic  info")
		}
	}

	return true
}


Try(hello, "ss", "ddd").Then(func(b bool) {
		fmt.Println(b)
}).Catch(func(err error) {
    fmt.Println(err, "err")
}).Finally(func(err *_KErr) {
    err.P()
})

Try(hello, "ss", "ddd", "a").Then(func(b bool) {
    fmt.Println(b)
}).Catch(func(err error) {
    fmt.Println(err, "err")
}).Finally(func(err *_KErr) {
    err.P()
})

Try(hello, "ss", "ddd", "a").Then(func(b bool) {
    fmt.Println(b)
}).Expect("sss %s", "ss")

```
