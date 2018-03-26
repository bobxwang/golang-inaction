> unsafe是不安全的，所以我们应该尽可能少的使用它，比如内存的操纵，这是绕过Go本身设计的安全机制的，不当的操作，可能会破坏一块内存，而且这种问题非常不好定位。
>
> 当然必须的时候我们可以使用它，比如底层类型相同的数组之间的转换；比如使用sync/atomic包中的一些函数时；还有访问Struct的私有字段时；该用还是要用，不过一定要慎之又慎。
>
> 还有，整个unsafe包都是用于Go编译器的，不用运行时，在我们编译的时候，Go编译器已经把他们都处理了。

##### unsafe.Pointer

> 是一种特殊意义的指针,可以包含任意类型的地址,有点类似于 C 语言里的 void* 指针

```go
i:= 10
ip:=&i
var fp *float64 = (*float64)(ip) // 这样会提示 cannot convert ip (type *int) to type *float64, 也就是不能进行强制转换,我们需要 unsafe.Pointer 

var fp *float64 = (*float64)(unsafe.Pointer(ip))
*fp = *fp * 3
fmt.Println(i) // 可以发现 i 的值也更改为了 30 
```

###### 使用四个原则

> *T 是不能计算偏移量的,也不能进行计算,但是 uintptr 可以, 所以这就是第三/四两个原则,转换后我们就可以进行特定的内存,达到对不同的内存读写目的 

* 任何指针都可以转换为 unsafe.Pointer
* unsafe.Pointer 可以转换为任何指针 
* unitptr 可以转换为 unsafe.Pointer
* unsafe.Pointer 可以转换为 uintptr  

```go
func main() {
	u:=new(user)
	fmt.Println(*u)
	pName:=(*string)(unsafe.Pointer(u))
	*pName="张三"
	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)))
	*pAge = 20
	fmt.Println(*u)
}
type user struct {
	name string
	age int
}
```

以上代码我们通过内存偏移的方式,定位到我们需要操作的字段,然后改变他们的值

我们看到第二个偏移的表达式非常长,有些同学可能会分段,做成下面这样

```go
temp := uintptr(unsafe.Pointer(u) + unsafe.Offsetof(u.age))
pAge := (*int)(unsafe.Pointer(temp))
*pAge = 20
```

逻辑上看,以上代码不会有问题,但这里会涉及到 GC, 如果我们的这些临时变量被 GC, 那么导致的内存操作就错了, 操作的也不知道是哪块内存了, 会引起莫名其妙的问题 