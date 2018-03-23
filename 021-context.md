> 控制并发有两种经典方式, 一种就是前面所讲的 WatiGroup, 另外一种就是今天的主角 Context 

* WaitGroup

  > WaitGroup以前我们在并发的时候介绍过，它是一种控制并发的方式，它的这种方式是控制多个goroutine同时完成

  ```go
  func main() {
  	var wg sync.WaitGroup
  	wg.Add(2)
  	go func() {
  		time.Sleep(2*time.Second)
  		fmt.Println("1号完成")
  		wg.Done()
  	}()
  	go func() {
  		time.Sleep(2*time.Second)
  		fmt.Println("2号完成")
  		wg.Done()
  	}()
  	wg.Wait()
  	fmt.Println("好了，大家都干完了，放工")
  } 
  ```

  这种方式适用于好多个 goroutine 协同做一件事情的时候,因为每个 goroutine 做的都是事情的一部分, 只有全部完成这事才算完成, 就像 Java 并发包中的栅栏一样 

* Context

  > 我们都知道一个goroutine启动后，我们是无法控制他的，大部分情况是等待它自己结束，那么如果这个goroutine是一个不会自己结束的后台goroutine呢？比如监控等，会一直运行的, 如果需要停止就要我们在外面通知它. 又或者一个网络请求 Request, 每个Request 又开启一个 goroutine 做些事情 , 可能这些 goroutine 又会开启其它的 goroutine, 所以我们需要一种可以跟跑完 goroutine的方案, 这就是 Context, 我们称之为上下文  

  ```go
  package main 

  import "fmt"
  import "context"
  import "time"

  func main() {
      ctx, cancel := context.WithCancel(context.Background())
  	go watch(ctx,"【监控1】")
  	go watch(ctx,"【监控2】")
  	go watch(ctx,"【监控3】")
  	time.Sleep(10 * time.Second)
  	fmt.Println("可以了，通知监控停止")
  	cancel()
  	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
  	time.Sleep(5 * time.Second)
  }
  func watch(ctx context.Context, name string) {
  	for {
  		select {
  			case <- ctx.Done():
  				fmt.Println(name,"退出监控...")
  				return 
  			default:
  				fmt.Println(name,"goroutine在工作中...")
  				time.Sleep(2 * time.Second)
  		}
  	}
  }
  ```

##### Context 接口定义 

```go
type Context interface {
  	Deadline() (deadline time.Time, ok bool) // 第一个参数为截止时间,到这个时间点会自动发起取消请求,第二个参数为false表明没有设置截止时间,如果需要取消要你自己主动调用取消函数进行取消
  	Done() <-chan struct {} // 如果该方法返回的chan可以读取意味着已经发起了取消请求,此时应该做清理释放动作来退出 goroutine 
  	Err() error // 返回取消的错误原因,因为什么Context被取消 
  	Value(key interface{}) interface{} // Context绑定的值, 是一个键值对,线程安全 
}
```

Context接口并不需要我们实现，Go内置已经帮我们实现了2个，我们代码中最开始都是以这两个内置的作为最顶层的partent context，衍生出更多的子Context

```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)
func Background() Context {
	return background
}
func TODO() Context {
	return todo
}
```

##### Context 继承衍生

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

下面我们演示通过WithValue来传递元数据 

```go
var key string = "name"
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx,key,"【监控1】") //附加值
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key),"监控退出，停止了...") //取出值
			return
		default:
			fmt.Println(ctx.Value(key),"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
```

##### Context 使用原则 

* 不要把Context放在结构体中，要以参数的方式传递
* 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位
* 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
* Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
* Context是线城安全的，可以放心的在多个goroutine中传递