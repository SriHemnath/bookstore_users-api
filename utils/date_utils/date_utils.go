package date_utils

import (
	"time"
)

const (
	standardDateLayout = "02-01-2006T15:04:05Z" //DDMMYYY
	dbDateLayout       = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC() //setting standard utc time for all timezones
}

func GetNowString() string {
	return time.Now().UTC().Format(standardDateLayout)
}

func GetNowDBFormat() string {
	return time.Now().UTC().Format(dbDateLayout)
}
