package main 

import (
    "fmt"
    "./pack1"
)

func main() {
    var test1 string
    test1 = pack1.ReturnStr()
    fmt.Printf("ReturnStr form pack1: %s\n", test1)
    fmt.Printf("Integer from pack1: %d\n", pack1.Pack1Int)
}