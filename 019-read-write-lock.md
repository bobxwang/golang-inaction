> 在 [goroutine](017-goroutine.md) 中讲资源竞争时提到了互斥锁, 其实很多时候我们不需要做读读控制, 因为读取并不是问题,问题主要是修改,修改的数据要同步,这样其他goroutine才可以感知到,所以真正的互斥应该是读取和修改,修改和修改之间,读取和读取是没有互斥操作的

#### 读写锁

> 读写锁可以让多个读操作同时并发,同时读取,但是对于写操作是完全互斥的.也就是说,当一个goroutine进行写操作的时候,其他goroutine既不能进行读操作,也不能进行写操作

##### 代码示例

```go
package main

import (
	"fmt"
  	"sync"
)

var count int 
var wg sync.WaitGroup
var rw sync.RWMutex  // 定义读写锁 

func main() {
	wg.Add(10)
	for i:=0;i<5;i++ {
		go read(i)
	}
	for i:=0;i<5;i++ {
		go write(i);
	}
	wg.Wait()
}
func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n",n)
	v := count
	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n,v)
	wg.Done()
	rw.RUnlock()
}
func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n",n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n,v)
	wg.Done()
	rw.Unlock()
}
```

我们在 `read` 里使用读锁，也就是 `RLock` 和 `RUnlock  `写锁的方法名和我们平时使用的一样 `Lock` 和 `Unlock` ,这样，我们就使用了读写锁，可以并发的读，但是同时只能有一个写，并且写的时候不能进行读操作，现在我们再运行代码，可以从输出的数据看到，可以独到新值了

我们同时也可以使用 `go build -race` 检测，也没有竞争提示了