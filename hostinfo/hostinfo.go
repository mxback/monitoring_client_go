package hostinfo

import (
	"github.com/shirou/gopsutil/host"
	"monitoring_client/util"
	"strconv"
)

type HostInfo struct {
	Name     string
	Update   int
	BootTime int
}

func GetHost() (hostInfo *HostInfo, err error) {
	host, err := host.Info()
	if err != nil {
		return
	}
	//fmt.Println(host.Hostname, host.Uptime/util.HourNum, host.BootTime/util.DayNum)
	hostInfo = &HostInfo{
		Name:     host.Hostname,
		Update:   int(host.Uptime / util.HourNum),
		BootTime: int(host.BootTime / util.DayNum),
	}
	return
}

func (h *HostInfo) GetMapHostInfo() (hostMap map[string]interface{}) {
	hostMapOne := map[string]string{}
	hostMapOne["hostname"] = h.Name
	hostMapOne["update"] = strconv.Itoa(h.Update)
	hostMapOne["boottime"] = strconv.Itoa(h.BootTime)
	hostMap = make(map[string]interface{})
	hostMap["hostinfo"] = hostMapOne
	return
}
