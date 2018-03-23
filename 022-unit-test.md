> 单元测试一般用来测试我们的代码逻辑有没有问题,有没有按照我们期望的运行,以保证代码质量 

##### main.go

``` go
func Add(a,b int) int {
  return a + b 
}
```

##### main-test.go

```go
func TestAdd(t *testing.T) {
	sum := Add(1,2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}
```

然后我们在终端项目目录下运行 `go test -v` 就可以看到测试结果了

```
hello go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
        main_test.go:26: the result is ok
PASS
```

Go语言为我们提供了测试框架,以便帮助我们更容易的进行单元测试,使用需遵循如下几点规则 

* 含有单元测试代码的go文件必须以`_test.go`结尾，Go语言测试工具只认符合这个规则的文件
* 单元测试文件名`_test.go`前面的部分最好是被测试的方法所在go文件的文件名，比如例子中是`main_test.go`，因为测试的`Add`函数，在`main.go`文件里
* 单元测试的函数名必须以`Test`开头，是可导出公开的函数
* 测试函数的签名必须接收一个指向`testing.T`类型的指针，并且不能返回任何值
* 函数名最好是Test+要测试的方法函数名，比如例子中是`TestAdd`，表示测试的是`Add`这个这个函数

遵循以上规则，我们就可以很容易的编写单元测试了，单元测试的重点在于测试代码的逻辑，场景等，以便尽可能的测试全面，保障代码质量逻辑。

#### 表组测试 

> 这个和基本的单元测试非常相似，只不过它是有好几个不同的输入以及输出组成的一组单元测试。

#### 模拟调用 

> 所测试方法不受依赖环境影响

#### 覆盖率

还是 `go test` 工具, 添加一个参数 `-coverprofile` 所以完整命令是 

```go
go test -v -coverprofile=/usr/tmp/go.out
```

-coverprofile 是指定生成的覆盖率文件

