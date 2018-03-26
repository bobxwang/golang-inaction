> 各其它语言一样, GO也有反射,这为我们提供了一种可以在运行时操作任意对象的能力

##### TypeOf/ValueOf

> 标准库为我们提供两种类型来分别表示他们 `reflect.Value` 和 `reflect.Type` ，并且提供了两个函数来获取任意对象的 `Value` 和 `Type` 

##### 动态调用方法

```go
func main() {
  	u := User{"李四",20}
  	v := reflect.ValueOf(u)
  	mPrint := v.MethodByName("Print")
  	if(mPrint.IsValie()) { // 使用IsValid 来判断是否可用（存在）
  		args := []reflect.Value{reflect.ValueOf("hello")}
  		fmt.Println(mPrint.Call(args))
    }
}
type User struct {
	Name string
  	Age int
}
func (u User) Print(prifx string) {
  	fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}
```

