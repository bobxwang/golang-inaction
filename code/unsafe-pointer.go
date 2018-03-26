package main 

import "fmt"
import "unsafe" 

func main() {
	
	i := 10
	ip := &i 

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3

	fmt.Println(i) // 30
}