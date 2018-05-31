//https://blog.csdn.net/sanxiaxugang/article/details/60871642
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var t time.Time
	t = time.Now()
	fmt.Println(t)
	fmt.Println(t.Hour())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Unix() * 1000)
	fmt.Println(strconv.FormatInt(t.UTC().UnixNano(), 10))

	var dur time.Duration
	dur = time.Duration(1527491550000)
	fmt.Printf("Nanoseconds: %v  Second:%v  Minutes: %v \n", dur.Nanoseconds(), dur.Seconds(), dur.Seconds())

}
