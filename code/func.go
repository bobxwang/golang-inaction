package main 

import (
    "fmt"
)

func main() {
    callback(1,Add) // it will display: The sum of 1 and 2 is: 3

    func() {
        sum := 0
        for i:=1;i<=10;i++ {
            sum += i
        }
        fmt.Printf("the sum value is: %d\n", sum)
    }() // 闭包调用
}

func Add(a, b int) {
    fmt.Printf("The sum of %d and %d is: %d\n", a, b, a + b)
}

func callback(y int, f func(int,int)) {
    f(y,2) // this becomes Add(y,2)
}