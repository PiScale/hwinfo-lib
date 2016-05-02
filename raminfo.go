package hwinfo

import (
	"os/exec"
	"strings"
)

type Ramtype struct {
	Size         string
	BankLocator  string
	ClockSpeed   string
	Manufacturer string
	SerialNumber string
	PartNumber   string
}

type Ramstats struct {
	TotalCapacity string
	Items         []Ramtype
}

//var Ram Ramstats

func Get_ram() (Ram Ramstats, err error) {
	cmd := exec.Command("/usr/sbin/dmidecode", "-t", "memory")
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println(err)
		return
	}
	output := strings.Split(string(buf), "\n")

	slot := -1

Iterate_Line:
	for _, line := range output {
		line_arr := strings.Split(line, ":")
		if len(line_arr) == 2 {
			key, value := line_arr[0], line_arr[1]
			if key == "\tMaximum Capacity" {
				Ram.TotalCapacity = value
				/* TODO: it's maximum(-able) ram capacity, not current total capacity.
				   Will fix in the future */
			}
			switch key {
			case "\tArray Handle":
				slot++
				var dimm Ramtype
				Ram.Items = append(Ram.Items, dimm)
			case "\tSize":
				Ram.Items[slot].Size = value
			case "\tBank Locator":
				Ram.Items[slot].BankLocator = value
			case "\tSpeed":
				Ram.Items[slot].ClockSpeed = value
			case "\tManufacturer":
				Ram.Items[slot].Manufacturer = value
			case "\tSerial Number":
				Ram.Items[slot].SerialNumber = value
			case "\tPart Number":
				Ram.Items[slot].PartNumber = value
			}

		} else {
			continue Iterate_Line
			/* Now it's just experimental "Label" in Go,
			because i changed my mind & use a different algo. */
		}

	}
	return
}
