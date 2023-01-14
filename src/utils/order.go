package utils

import (
	"fmt"
	"time"
	"winter-examination/src/conf"
)

var orderIdGenerator = make(chan string, 1)
var orderIdRequester = make(chan string, 1)
var a = map[string]int64{}

func InitOrderIdGenerator() {
	go OrderIdGenerateRoutine()
}
func OrderIdGenerateRoutine() {
	for {
		select {
		case <-orderIdRequester:
			DurationTimeStamp := time.Now().Unix() - conf.OrderIdBaseTimeStamp
			a[time.Now().Format("20060102")]++
			num := a[time.Now().Format("20060102")]
			sprintf := fmt.Sprintf("%d", DurationTimeStamp<<conf.OrderIdLeftShiftNumber|num)
			orderIdGenerator <- sprintf
		}
	}
}

func GetOrderId() string {
	orderIdRequester <- ""
	select {
	case id := <-orderIdGenerator:
		return id
	}
}
