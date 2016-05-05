package hwinfo

// Tuong tu nhu check NIC info, nhung voi /sys/block
/* # hdparm -iI /dev/sda
        Model Number:       Micron_M600_MTFDDAV256MBF
        Serial Number:      15090F0319BE
        Firmware Revision:  MA01
        Form Factor: less than 1.8 inch
        Nominal Media Rotation Rate: Solid State Device
        device size with M = 1024*1024:      244198 MBytes
		device size with M = 1000*1000:      256060 MBytes (256 GB)

   # lshw -class disk
       product: Micron_M600_MTFD
       bus info: scsi@0:0.0.0
       logical name: /dev/sda
       version: MA01
       serial: 15090F0319BE
       size: 238GiB (256GB)

Example dest symlink: /sys/devices/pci0000:00/0000:00:1f.2/ata1/host0/target0:0:0/0:0:0:0/block/sda
--> Base dir for information: /sys/devices/pci0000:00/0000:00:1f.2/ata1/host0/target0:0:0/0:0:0:0/*
*/

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	utils "github.com/PiScale/hwinfo-lib/utils"
)

type HddType struct {
	Model        string
	SerialNumber string
	Firmware     string
	Size         string
	DevName      string
	//Type         string  // currently, i don't have precise method to determine SSD, SAS, SATA,..
	Bus string // ide, ata, scsi
}

type HddStats struct {
	Items []HddType
}

func Get_hdd() (Hdd HddStats, err error) {
	baseHddpath := "/sys/block/"
	cmd := exec.Command("ls", baseHddpath)
	buf, err := cmd.Output()
	if err != nil {
		return
	}

	output := string(buf)
	for _, device := range strings.Split(output, "\n") {
		if len(device) > 1 {
			devPath := baseHddpath + device
			destLink, err := filepath.EvalSymlinks(devPath)
			if err == nil {
				if strings.Contains(destLink, "pci") { // It's physical disk
					disk, err := get_disk_info(destLink)
					if err == nil {
						Hdd.Items = append(Hdd.Items, disk)

					}
				}
			}
		}
	}

	return
}

func get_disk_info(destLink string) (info HddType, err error) {
	//fmt.Printf("%v has type %T\n", destLink, destLink)
	destLinkArr := strings.Split(destLink, "/block/")
	infoDir := destLinkArr[0]
	info.DevName = destLinkArr[1]
	//fmt.Println("infoDir =", infoDir)

	info.Model, _ = utils.Cat_file(infoDir + "/model")
	info.Firmware, _ = utils.Cat_file(infoDir + "/rev")

	devPath := "/dev/" + info.DevName
	//fmt.Println("devPath =", devPath)
	cmd := exec.Command("sfdisk", "-s", devPath)
	buf, err := cmd.Output()
	if err != nil {
		return
	}

	output, _ := strconv.Atoi(strings.Split(string(buf), "\n")[0]) // unit: KB
	switch {
	case output > 1000*1000*1000: // TB
		size := 0.1 + float64(output)/(1000*1000*1000) // 0.1 is a trick because of miserly vendors!
		info.Size = fmt.Sprintf("%.2f", size) + " TB"
	case output > 1000*1000: // GB
		size := 0.1 + float64(output)/(1000*1000)
		info.Size = fmt.Sprintf("%.2f", size) + " GB"
	}

	/* Sadly, i found this command below later.
	Therefore i keep above codes for reference and knowledge. */
	cmd = exec.Command("/sbin/udevadm", "info", "--query=property", "--name="+info.DevName)
	buf, err2 := cmd.Output()
	if err2 != nil {
		return
	}

	for _, line := range strings.Split(string(buf), "\n") {
		lineArr := strings.Split(line, "=")
		if len(lineArr) == 2 {
			key, value := lineArr[0], lineArr[1]
			switch key {
			case "ID_BUS":
				info.Bus = value
			case "ID_MODEL":
				info.Model = value
			case "ID_SERIAL_SHORT":
				info.SerialNumber = value
			case "ID_REVISION":
				info.Firmware = value

			}
		}
	}
	//fmt.Println(info)
	return
}
