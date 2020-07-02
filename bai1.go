package main

import (
	"fmt"
	"time"
)

func Print() {
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		t := time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println("time now: ", t)

	}
	fmt.Println("kết thúc")
}
