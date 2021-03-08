package main

import (
	"encoding/json"
	"fmt"
	"monitoring_client/db"
	"monitoring_client/device"
	"time"
)

func main() {
	redisDb := &db.RedisInfo{
		Address:  "192.168.205.130:6379",
		PassWord: "",
		Zone:     13,
	}
	err := redisDb.Init()
	if err != nil {
		fmt.Println("redis conn fail")
	}
	defer redisDb.Client.Close()
	for {
		deviceInfo, err := device.GetDeviceInfo()
		if err != nil {
			fmt.Println(err)
		}
		deviceJOSN, err := json.Marshal(deviceInfo.DeviceInfoList)
		deviceString := string(deviceJOSN)
		_, err = redisDb.Client.Do("Set", deviceInfo.IpAddress.IP, deviceString)
		if err != nil {
			return
		}
		time.Sleep(time.Second * 10)
	}
}
