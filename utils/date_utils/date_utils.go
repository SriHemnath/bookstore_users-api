package date_utils

import (
	"time"
)

const (
	apiDateLayout = "02-01-2006T15:04:05Z" //DDMMYYY
)

func GetNow() time.Time {
	return time.Now().UTC() //setting standard utc time for all timezones
}

func GetNowString() string {
	return time.Now().UTC().Format(apiDateLayout)
}
