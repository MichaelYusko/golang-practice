package main

import (
	"fmt"
	"golang-practice/models"
	"time"
)

func main() {
	user := models.NewUser("Mike", 150.00, "Tokyo")
	alarm := models.NewOwnAlarm(1, "Hello")
	fmt.Println(user)
	startAlarm(alarm)
}

func startAlarm(alarm models.Alarm) {
	alarm.SetTime(time.Now())
	alarm.SetAlarmSound("wake upp!")
	alarm.Play()
}
