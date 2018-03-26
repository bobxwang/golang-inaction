package main

import "fmt"

/* 定义结构体 */
type Circle struct {
	radius float64
}

/* 全局变量 */
var global_var int 

func main() {
    fmt.Println("hello, world")

    var init, added = 69000, 76000
    fmt.Printf("共投入了 %d 元\n", init + added)
    fmt.Printf("最大值是 %d\n", max(init,added))

    a, b := sswap("world", "hello")
    fmt.Println(a,b)

    var c int = 100
   	var d int= 200
    fmt.Printf("交换前，c 的值 : %d\n", c )
    fmt.Printf("交换前，d 的值 : %d\n", d )
    swap(&c, &d)
    fmt.Printf("交换后，c 的值 : %d\n", c )
    fmt.Printf("交换后，d 的值 : %d\n", d )

    nextNumber := getSequence()
    /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   	fmt.Println(nextNumber())
   	fmt.Println(nextNumber())
   	fmt.Println(nextNumber())
   	/* 创建新的函数 nextNumber1，并查看结果 */
   	nextNumber1 := getSequence()  
   	fmt.Println(nextNumber1())
   	fmt.Println(nextNumber1())

   	var c1 Circle
   	c1.radius = 10.00
   	fmt.Println("Area of Circle(c1) = ", c1.getArea())

   	var numbers = make([]int,3,5)
    printSlice(numbers)

    var phone Phone
    phone = new(NokiePhone)
    phone.call()

    // 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
        fmt.Println("100/10 = ", result)
    }
    // 当被除数为零的时候会返回错误信息
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }

    sum(1,2)
    sum(1,2,3,4,5)
    nums := []int{1,2,3,4,5,6,7,8,9,10}
    sum(nums...)
}

/* 求最大值 */
func max(num1 int, num2 int) int {
	var rs int

	if(num1 > num2) {
		rs = num1
	} else {
		rs = num2
	}

	return rs 
}

/* 交换两个值位置返回,可以返回类似元组的东西 */
func sswap(x,y string) (string,string) {
	return y, x	
}

/* 引用传递,而不是值传递 */
func swap(x,y *int) {
	var temp int 
	temp = *x
	*x = *y
	*y = temp 
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i 
	}
}

/* 这样定义后就相当于 getArea 这方法是 Circle 类型对象中的方法, 相当于C#中的扩展方法 */
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func printSlice(x []int) {
  fmt.Printf("len=%d cap=%d slice=%v\n", len(x),cap(x), x)
}

/* 定义一个接口 */
type Phone interface {
	call()
}

type NokiePhone struct { }
func (NokiePhone NokiePhone) call() {
	fmt.Println("I am NokiePhone, I can call u ")
}

type DivideError struct {
  	dividee int
  	divider int 
}

// 实现 `error` 接口
func (de DivideError) Error() string {
  	strFormat := `
    	Cannot proceed, the divider is zero.
    	dividee: %d
    	divider: 0`
    return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee: varDividee,
            divider: varDivider,
        }
        errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider, ""
    }
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
