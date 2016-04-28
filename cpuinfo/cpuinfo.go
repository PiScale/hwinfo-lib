package cpuinfo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Cpustats struct {
	Model      string
	Quantity   int
	Totalcores int
}

/*
Read file /proc/cpuinfo, then parse it line by line. Read line have info we need
and bypass line have info we don't need. Count CPU quantity and vcores number.
*/
const CPU_INFO_FILE = "/proc/cpuinfo"

var Cpu Cpustats

func init() {
	b, err := ioutil.ReadFile(CPU_INFO_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpulines := strings.Split(string(b), "\n")
	for _, line := range cpulines {
		line_arr := strings.Split(line, ":")
		if len(line_arr) == 2 {
			key, value := line_arr[0], line_arr[1]
			switch key {
			case "model name\t":
				Cpu.Model = value
			case "processor\t":
				Cpu.Totalcores++
			case "physical id\t":
				if cpuId, err := strconv.Atoi(string(value)); err != nil && Cpu.Quantity == cpuId {
					Cpu.Quantity++
				}
			}
		}
	}
}
