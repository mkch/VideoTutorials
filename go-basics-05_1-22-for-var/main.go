package main

import (
	"log"
	"time"
	"unsafe"
)

func main() {
	for i := 0; i < 3; i++ {
		DebugPrint("main", i, &i)
		go func() {
			time.Sleep(time.Millisecond * 1000)
			DebugPrint("go", i, &i)
		}()
	}

	time.Sleep(time.Second * 3)
}

// DebugPrint prints v and p with log.Printf without escaping them.
func DebugPrint(prefix string, v int, p *int) {
	log.Printf("%s: %v 0x%x", prefix, v, uintptr(unsafe.Pointer(p)))
}
