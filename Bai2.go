package main

import "fmt"

func GetDay(a int64) {
	day := a / 86400
	fmt.Println("So ngay tinh tu 01/01/1970 den nay: ", day)
}
