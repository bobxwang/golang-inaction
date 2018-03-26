package main 

import "log"

func init(){
	log.SetPrefix("【UserCenter】") // 设置日志前缀,可以是业务名称 
	log.SetFlags(log.Ldate|log.Lshortfile|log.Ltime|log.LUTC)
}

func main() {
	log.Println("aaaa")
}

/*
除了Println,还有Fatal以及Panic系列的函数
其中Fatal表示程序遇到了致命的错误，需要退出，这时候使用Fatal记录日志后，然后程序退出，也就是说Fatal相当于先调用Print打印日志，然后再调用os.Exit(1)退出程序。
同理Panic系列的函数也一样，表示先使用Print记录日志，然后调用panic()函数抛出一个恐慌，这时候除非使用recover()函数，否则程序就会打印错误堆栈信息，然后程序终止。
*/