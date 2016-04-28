/* Read from /sys/devices/virtual/dmi/id/board_{name,serial,vendor} */
package mbinfo

import (
	"fmt"
	"os/exec"
)

const MB_NAME_FILE = "/sys/devices/virtual/dmi/id/board_name"
const MB_SERIAL_FILE = "/sys/devices/virtual/dmi/id/board_serial"
const MB_VENDOR_FILE = "/sys/devices/virtual/dmi/id/board_vendor"

type MBstats struct {
	Model  string
	Serial string
}

var Motherboard MBstats

func cat_file(filepath string) (ret string) {
	cmd := exec.Command("/bin/cat", filepath)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	ret = string(buf)
	return
}

func init() {
	Motherboard.Model = cat_file(MB_VENDOR_FILE) + cat_file(MB_NAME_FILE)
	Motherboard.Serial = cat_file(MB_SERIAL_FILE)
}
