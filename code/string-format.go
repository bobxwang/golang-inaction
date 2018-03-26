package main 

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	p := point{1,2}

	fmt.Printf("%v\n", p)  // %v 

	fmt.Printf("%+v\n", p) // 如果是个结构体, %+v的输出将包括结体体的字段名

	fmt.Printf("%#v\n", p) // %#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。

	fmt.Printf("%T\n", p)  // 需要打印值的类型，使用 %T。

	fmt.Printf("%t\n", true) 

	fmt.Printf("%d\n", 123) // 标准十进制类型使用 %d

	fmt.Printf("%b\n", 14)  // 这个输出二进制表示形式

	fmt.Printf("%c\n", 33)  // 这个输出给定整数的对应字符

	fmt.Printf("%x\n", 456)  // 十六进制编码 

	fmt.Printf("%f\n", 78.9) // 对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化

	fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)  // %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式

    fmt.Printf("%s\n", "\"string\"")  // 使用 %s 进行基本的字符串输出

    fmt.Printf("%q\n", "\"string\"")  // 像 Go 源代码中那样带有双引号的输出，使用 %q

 	fmt.Printf("%x\n", "hex this")  // 和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示

 	fmt.Printf("%p\n", &p)  // 要输出一个指针的值，使用 %p

 	// 当输出数字的时候，你将经常想要控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。
 	// 默认结果使用右对齐并且通过空格来填充空白部分
    fmt.Printf("|%6d|%6d|\n", 12, 345) 

    // 你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // 要左对齐，使用 - 标志

    // 你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // 要左对齐，和数字一样，使用 - 标志
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

/*
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0xc420012060
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error
*/