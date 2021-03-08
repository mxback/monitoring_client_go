package alldisk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"monitoring_client/util"
	"strconv"
)

type DiskList struct {
	DiskList []*Disk
}

type Disk struct {
	name  string
	total int
	used  int
}

func GetDisk() (diskList DiskList, err error) {

	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		//fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
		disk := &Disk{
			name:  part.Mountpoint,
			total: int(diskInfo.Total) / util.IntNum,
			used:  int(diskInfo.Used) / util.IntNum,
		}
		diskList.DiskList = append(diskList.DiskList, disk)
	}
	return
}

func (d *DiskList) GetMapDisk() (diskMap map[string]interface{}) {
	var diskMapTwo = map[string]map[string]string{}
	for _, value := range d.DiskList {
		diskMapOne := map[string]string{}
		diskMapOne["used"] = strconv.Itoa(value.used)
		diskMapOne["total"] = strconv.Itoa(value.total)
		diskMapTwo[value.name] = diskMapOne
	}
	diskMap = make(map[string]interface{})
	diskMap["disk"] = diskMapTwo
	return
}

/*
总结了golang中字符串和各种int类型之间的相互转换方式：

string转成int：
int, err := strconv.Atoi(string)
string转成int64：
int64, err := strconv.ParseInt(string, 10, 64)
int转成string：
string := strconv.Itoa(int)
int64转成string：
string := strconv.FormatInt(int64,10)
————————————————
版权声明：本文为CSDN博主「排骨瘦肉丁」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/iamlihongwei/article/details/79550958


	diskList, _ := alldisk.GetDisk()
	for _, diskInfo := range diskList.DiskList {
		fmt.Println(diskInfo.Name, diskInfo.Total, diskInfo.Used)
	}
*/
