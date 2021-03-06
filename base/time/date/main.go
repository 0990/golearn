package main

import (
	"time"
	"fmt"
)

const (
	goBornDate       = "2006-01-02"          //长日期格式
)
func main(){
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Minute)

	str1 := t1.Format("2006-01-02 15:04:05")
	str2 := t2.Format("2006-01-02 15:04:05")

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-02 15:04:05", time.Local)
	fmt.Println(t1, t1.Before(t2), str1, str2, t)
}
