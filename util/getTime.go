package util

import "time"

func GetTimeMap() (timeMap map[string]interface{}) {
	timeUnix := time.Now().Unix()
	timeMap = map[string]interface{}{
		"time": timeUnix,
	}
	return
}
