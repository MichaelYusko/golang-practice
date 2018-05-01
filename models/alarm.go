package models

import "time"

// Interface for the OwnAlarm structure
type Alarm interface {
	SetTime(time.Time)
	SetAlarmSound(string)
	Play()
}

type OwnAlarm struct {
	ID int
	WelcomeMessage string
}

// Create a new OwnAlarm
func NewOwnAlarm(id int, message string) *OwnAlarm {
	return &OwnAlarm{id, message}
}

func (a *OwnAlarm) SetTime(time time.Time) {
	panic("implement me!")
}

func (a *OwnAlarm) SetAlarmSound(sound string) {
	panic("implement me!")
}

func(a *OwnAlarm) Play() {
	panic("implement me!")
}