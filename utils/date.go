package utils

import (
	"time"
)

// GetCurrTime get expire timestamp
func GetCurrTime() int {
	return int(time.Now().Unix())
}
