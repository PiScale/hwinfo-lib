package main

import (
	"fmt"
	. "github.com/minhchuduc/hwinfo-lib/cpuinfo"
	. "github.com/minhchuduc/hwinfo-lib/mbinfo"
)

func main() {
	fmt.Println(Cpu)
	fmt.Println(Motherboard)
}
