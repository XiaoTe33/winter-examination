package utils

import (
	"fmt"
	"time"
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
			const BaseTimeStamp int64 = 1672531200
			DurationTimeStamp := time.Now().Unix() - BaseTimeStamp
			a[time.Now().Format("20060102")]++
			num := a[time.Now().Format("20060102")]
			sprintf := fmt.Sprintf("%d", DurationTimeStamp<<32|num)
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
