// Package datetime for getting date and time
package datetime

import "time"

var currentTime = time.Now()

// CurrentTime has complete time info
var CurrentTime = currentTime

// GetHour for getting current Hour
func GetHour() int {
	return currentTime.Hour()
}

// GetMinute for getting current Minute
func GetMinute() int {
	return currentTime.Minute()
}

// GetSecond for getting current Second
func GetSecond() int {
	return currentTime.Second()
}
