package main 

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
    "strings"
)

var (
    questiontotal string
    operationsymbol [2]string
)

func main() {
    fmt.Println("输入你要生成的题目个数，请输整数：")
    fmt.Scanln(&questiontotal)
    qt, error := strconv.Atoi(questiontotal)
    if error != nil {
        fmt.Println("输入有误，请重新输入")
    }
    
    rand.Seed(time.Now().Unix())
    fmt.Printf("你输入的数字是 %d，将生成 %d 道题目，请稍等\n", qt, qt)

    operationsymbol[0] = "+"
    operationsymbol[1] = "-"
    // 题目是Key,答案是Value 
    var operations map[string]string = make(map[string]string)

    for i:=0; i<qt; i++ {
        var ino1 = rand.Intn(100)
        var ino2 = rand.Intn(100)
        oper := operationsymbol[rand.Intn(2)]
        if(oper == "-") {
            if(ino1 < ino2) {
                // 防止结果是负数 
                temp := ino2
                ino2 = ino1
                ino1 = temp
            }
        }
        no1 := strconv.Itoa(ino1)
        no2 := strconv.Itoa(ino2)

        key := strings.Join([]string{no1,oper,no2,"="}," ")
        var value int 
        if(oper == "-") {
            value = ino1 - ino2
        } else {
            value = ino1 + ino2
        }

        operations[key] = strconv.Itoa(value)
    }

    fmt.Println("开始做题\n")
    //记录开始时间
    start := time.Now()
    var uinput string
    for key, value := range operations {
        fmt.Println(key)
        for {
            fmt.Scanln(&uinput)
            if (uinput == value) {
                fmt.Println("恭喜你，做对了，太棒了")
                break
            } else {
                fmt.Printf("你输入的答案是 %s，不过做错了，请重做\n", uinput)
            }
        }
    }

    //记录结束时间
    end := time.Now()
    //输出执行时间，单位为秒。
    totaltime := ( end.Sub(start).Nanoseconds() / 1000000 ) / 1000
    var averagetime = float32(totaltime * 1.0) / float32(qt * 1.0)
    fmt.Printf("总共做了 %d 道题目，共花费了 %d 秒，平均一道题目消耗了 %.2f 秒\n", qt, totaltime, averagetime)  

    fmt.Println()
}