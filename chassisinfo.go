package hwinfo

import (
	"os/exec"
	"strings"
)

type ChassisStats struct {
	Manufacturer string
	SerialNumber string
}

func Get_chassis() (Chassis ChassisStats, err error) {
	cmd := exec.Command("/usr/sbin/dmidecode", "-t", "chassis")
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println(err)
		return
	}

	output := strings.Split(string(buf), "\n")

	for _, line := range output {
		line_arr := strings.Split(line, ":")

		if len(line_arr) == 2 {
			key, value := line_arr[0], line_arr[1]
			switch key {
			case "\tManufacturer":
				Chassis.Manufacturer = value
			case "\tSerial Number":
				Chassis.SerialNumber = value

			}
		}
	}
	return
}
