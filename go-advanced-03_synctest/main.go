package main

import (
	"testing/synctest"
	"time"
)

func Business() {
	BusinessStep1()
	// 项目经理要求这里运行缓慢，好让客户给钱优化
	time.Sleep(time.Second * 2)
	BusinessStep2()
}

func main() {
	Business()
	synctest.Run(Business)
}
