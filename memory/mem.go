package memory

import (
	"github.com/shirou/gopsutil/mem"
	"monitoring_client/util"
	"strconv"
)

type Memory struct {
	Total int
	Used  int
}

func GetMem() (memory *Memory, err error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	total := int(memStat.Total / util.IntNum)
	used := int(memStat.Used / util.IntNum)
	memory = &Memory{
		Total: total,
		Used:  used,
	}
	return
}

func (m *Memory) GetMapMemory() (MemoryMap map[string]interface{}) {
	MemoryMapOne := map[string]string{}
	MemoryMapOne["usage"] = strconv.Itoa(m.Used)
	MemoryMapOne["total"] = strconv.Itoa(m.Total)
	MemoryMap = make(map[string]interface{})
	MemoryMap["memory"] = MemoryMapOne
	return
}

