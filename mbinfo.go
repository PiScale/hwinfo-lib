package hwinfo

import (
	utils "github.com/PiScale/hwinfo-lib/utils"
	//"os/exec"
	"strings"
)

/* Read from /sys/devices/virtual/dmi/id/board_{name,serial,vendor} */
const MB_NAME_FILE = "/sys/devices/virtual/dmi/id/board_name"
const MB_SERIAL_FILE = "/sys/devices/virtual/dmi/id/board_serial"
const MB_VENDOR_FILE = "/sys/devices/virtual/dmi/id/board_vendor"

type MBstats struct {
	Model        string
	SerialNumber string
}

var Motherboard MBstats

/* Moved to utils package
func Cat_file(filepath string) (ret string, err error) {
	cmd := exec.Command("/bin/cat", filepath)
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println(err)
		return
	}
	ret = string(buf)
	return
}
*/

func Get_motherboard() (Motherboard MBstats, err error) {
	vendor, _ := utils.Cat_file(MB_VENDOR_FILE)
	productName, _ := utils.Cat_file(MB_NAME_FILE)
	Motherboard.Model = vendor + productName
	Motherboard.Model = strings.Replace(Motherboard.Model, "\n", " ", -1)

	Motherboard.SerialNumber, err = utils.Cat_file(MB_SERIAL_FILE)
	if err == nil {
		Motherboard.SerialNumber = strings.Replace(Motherboard.SerialNumber, "\n", " ", -1)
	}

	return
}
