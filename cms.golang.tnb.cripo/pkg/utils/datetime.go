package utils

import (
	"time"
)

func ParseTime(value string) (time.Time, error) {
	// dthr, err := time.Parse("02/01/2006 15:04:05", value)
	// if err != nil {
	// 	return time.Time{}, err
	// }
	// return dthr, nil
	return time.Parse("02/01/2006 15:04:05", value)
}
