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
