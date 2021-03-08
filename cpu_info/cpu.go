package cpu_info

import (
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

type Cpu struct {
	Total int
	Usage int
}

func GetCpu() (cpuInfo *Cpu, err error) {
	_, err = cpu.Info()
	if err != nil {
		return
	}

	total, err := cpu.Counts(false)
	if err != nil {
		return
	}
	info, _ := cpu.Percent(time.Duration(time.Second), false)

	cpuInfo = &Cpu{
		Total: total,
		Usage: int(info[0]),
	}
	return
}

func (c *Cpu) GetMapCpu() (cpuMap map[string]interface{}) {
	cpuMapOne := map[string]string{}
	cpuMapOne["usage"] = strconv.Itoa(c.Usage)
	cpuMapOne["total"] = strconv.Itoa(c.Total)
	cpuMap = make(map[string]interface{})
	cpuMap["cpu"] = cpuMapOne
	return
}
