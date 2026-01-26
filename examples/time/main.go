package main

import (
	"fmt"
	"time"
)

func GetDateTimeNow() string {
	return time.Now().Format(time.DateTime)
}

func GetDateTimeNowUTC() string {
	return time.Now().UTC().Format(time.DateTime)
}

func main() {
	fmt.Println("time now:", GetDateTimeNow())
	fmt.Println("time now in UTC:", GetDateTimeNowUTC())
}
