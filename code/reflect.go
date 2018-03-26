package main 

import "fmt"
import "reflect"

type User struct {
	Name string
	Age int 
}

func main() {
	u := User{"李四", 25}
	t := reflect.TypeOf(u)
	fmt.Println(t)  // 会打印出 main.User 
	v := reflect.ValueOf(u)
	fmt.Println(v)  // 会打印出 {李四 25}
	// 对于上面两种情况,还可以通过格式化来打印出相同的东西 
	fmt.Printf("%T\n",u)
	fmt.Printf("%v\n",u)

	for i:=0;i<t.NumField();i++ {
		fmt.Println(t.Field(i).Name)
	}
	for i:=0;i<t.NumMethod() ;i++  {
		fmt.Println(t.Method(i).Name)
	}

	fmt.Println(t.Kind())  // 会打印出 struct
	fmt.Println(v.Kind())  // 会打印出 struct

	u1 := v.Interface().(User)
	fmt.Println(u1)
	t1:=v.Type()
	fmt.Println(t1)
}