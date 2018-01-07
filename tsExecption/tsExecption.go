package tsExecption

import (
//	"runtime/debug"
	"fmt"
)

func GoRoutineEnd(end string)() {
	fmt.Println("<<<<<<<<<<<<<<<", end)
//	if r := recover(); r != nil {
//		fmt.Println(r)
//		debug.PrintStack()
//	}
	fmt.Println(">>>>>>>>>>>>>>>", end)
}