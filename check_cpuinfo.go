package main

import (
	"fmt"
	. "github.com/PiScale/hwinfo-lib/cpuinfo"
	. "github.com/PiScale/hwinfo-lib/mbinfo"
)

func main() {
	fmt.Println(Cpu)
	fmt.Println(Motherboard)
}
