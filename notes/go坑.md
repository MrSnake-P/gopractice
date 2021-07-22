# 接口的陷阱

```go
type People interface {
    Show()
}
type Student struct{}

func (stu *Student) Show() {
}

func main() {
    var s *Student
    if s == nil {
        fmt.Println("s is nil")
    } else {
        fmt.Println("s is not nil")
    }
    var p People = s
    if p == nil {
        fmt.Println("p is nil")
    } else {
        fmt.Println("p is not nil")
    }
}
___________________________________________
s is nil
p is not ni
```

> 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil
>
> p 的动态值是 nil，但是动态类型却是 `*Student`，是一个 nil 指针

