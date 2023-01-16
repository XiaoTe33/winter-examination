package utils

import (
	"fmt"
	"time"
	"winter-examination/src/conf"
)

var orderIdLock = make(chan struct{}, 1)
var a = map[string]int64{}

func GetOrderId() string {
	orderIdLock <- struct{}{}
	DurationTimeStamp := time.Now().Unix() - conf.OrderIdBaseTimeStamp
	a[time.Now().Format("20060102")]++
	num := a[time.Now().Format("20060102")]
	sprintf := fmt.Sprintf("%d", DurationTimeStamp<<conf.OrderIdLeftShiftNumber|num)
	<-orderIdLock
	return sprintf
}
