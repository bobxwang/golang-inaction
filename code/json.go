package main 

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

	// 以下是基本数据类型到JSON字符串的编码过程 
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // 下面是切片和MAP编码成JSON数组和对象的例子 
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // JSON 包可以自动的编码你的自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。在上面 Response2 的定义可以作为这个标签这个的一个例子。
    res2D := Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]interface{}
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
    num := dat["num"].(float64)
    fmt.Println(num)
	strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // 我们也可以解码 JSON 值到自定义类型。
    // 这个功能的好处就是可以为我们的程序带来额外的类型安全加强，并且消除在访问数据时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}

/*
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
&{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
*/