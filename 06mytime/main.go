package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006")) // DD-MM-YYYY
	fmt.Println(presentTime.Format("01-02-2006")) // MM-DD-YYYY
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("02-01-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	nextDay := time.Now().AddDate(0, 0, 1) //1 day future date
	fmt.Println("nextday::", nextDay)
}
