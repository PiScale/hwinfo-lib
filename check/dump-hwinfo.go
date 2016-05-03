package main

import (
	"fmt"
	hwinfo "github.com/PiScale/hwinfo-lib"
)

func main() {
	Cpu, _ := hwinfo.Get_cpu()
	Motherboard, _ := hwinfo.Get_motherboard()
	Ram, _ := hwinfo.Get_ram()
	Chassis, _ := hwinfo.Get_chassis()
	Nic, _ := hwinfo.Get_nic()

	fmt.Println(Cpu)
	fmt.Println(Motherboard)
	fmt.Println(Ram)
	fmt.Println(Chassis)
	fmt.Println(Nic)
}
