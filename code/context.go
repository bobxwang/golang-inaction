package main 

import "fmt"
import "context"
import "time"

var key = "hello"

func main() {

	//control_onegoroutine()

	//control_moregoroutine()

	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx,key,"带KEY的监控")
	go watchWithKey(valueCtx)
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

func watchWithKey(ctx context.Context) {
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

/*
启动三个goroutine, 每一个都使用了Context进行跟踪,当使用cancel通知取消时,这三个goroutine都会被取消
*/
func control_moregoroutine() {
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

func control_onegoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("退出监控")
				return 
			default:
				fmt.Println("goroutine在工作中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel() // 就是上面的cancel变量, 它是一个 CancelFunc 类型, 我们调用它来发出取消指令 
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}