## functions

```go
func Fatal(v ...any)
func Fatalf(format string, v ...any)
func Fatalln(... any)
```

先输入日志，再调用 os.Exit(1) 终结程序。

```go
func Flags() int 
```

返回当前日志的输出格式

