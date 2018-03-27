package main 

import "fmt"
import "unsafe" 

func main() {
	
	i := 10
	ip := &i 

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3

	fmt.Println(i) // 30

	fun1()
	fun2()
	fun3()
}

func fun1()  {
	/*
	unsafe.Pointer(&a)是一个原生指针,指向的内存保存了一个int 
	c := (*string)(unsafe.Pointer(&a)) 将原生指针转换为了指向string的指针, 但是c指向的内存只分配了一个int的长度 
	*/
    a := 2
    c := (*string) (unsafe.Pointer(&a))
    *c = "44"
    fmt.Println(*c) // 空白
}

func fun2()  {
    a := "654"
    c := (*string) (unsafe.Pointer(&a))
    *c = "44"
    fmt.Println(*c) // 打印出 44  
}

func fun3()  {
    a := 3
    c := *(*string) (unsafe.Pointer(&a))
    c = "445"
    fmt.Println(c) // 打印出 445 
}