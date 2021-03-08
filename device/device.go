package device

import (
	"monitoring_client/alldisk"
	"monitoring_client/cpu_info"
	"monitoring_client/hostinfo"
	"monitoring_client/ip_address"
	"monitoring_client/memory"
	"monitoring_client/util"
)

type DeviceInfo struct {
	CPU            *cpu_info.Cpu
	Memory         *memory.Memory
	HostInfo       *hostinfo.HostInfo
	DiskList       alldisk.DiskList
	IpAddress      *ip_address.IpAddress
	DeviceInfoList []map[string]interface{}
}

func GetDeviceInfo() (deviceInfo *DeviceInfo, err error) {
	diskList, err := alldisk.GetDisk()
	if err != nil {
		return
	}

	cpuInfo, err := cpu_info.GetCpu()
	if err != nil {
		return
	}

	hostInfo, err := hostinfo.GetHost()
	if err != nil {
		return
	}

	memoryInfo, err := memory.GetMem()
	if err != nil {
		return
	}

	ipAddress, err := ip_address.GetOutboundIP()
	if err != nil {
		return
	}
	deviceInfo = &DeviceInfo{
		CPU:       cpuInfo,
		DiskList:  diskList,
		HostInfo:  hostInfo,
		Memory:    memoryInfo,
		IpAddress: ipAddress,
	}

	deviceInfo.DeviceInfoList = append(deviceInfo.DeviceInfoList, deviceInfo.CPU.GetMapCpu(),
		deviceInfo.DiskList.GetMapDisk(), deviceInfo.HostInfo.GetMapHostInfo(), deviceInfo.Memory.GetMapMemory(),
		util.GetTimeMap())
	return
}

func (d *DeviceInfo) GetDeviceInfoMap() (deviceInfoMap map[string]interface{}) {
	deviceInfoMap = make(map[string]interface{})
	if d.IpAddress.IP != "" || d.DeviceInfoList != nil {
		deviceInfoMap[d.IpAddress.IP] = d.DeviceInfoList
	}
	return
}
