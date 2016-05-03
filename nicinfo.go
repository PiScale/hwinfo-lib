package hwinfo

/*
Find all network interface: "ls /sys/class/net/"
Check destination of each result (symlink) above.
If it's have "pci" in destination --> it's a physical interface.
With each physical interface --> find Vendor & MAC_addr of it.
*/

// Should use filepath.EvalSymlinks
// Ref: http://stackoverflow.com/questions/18062026/resolve-symlinks-in-go

import (
	"os/exec"
	"path/filepath"
	"strings"

	utils "github.com/PiScale/hwinfo-lib/utils"
)

type NicType struct {
	IfName string
	MAC    string
	//Vendor   string  // TODO: Should add in the future.
	//Product  string
}

type NicStats struct {
	Items []NicType
}

func Get_nic() (Nic NicStats, err error) {
	baseNicPath := "/sys/class/net/"
	cmd := exec.Command("ls", baseNicPath)
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println("Error:", err)
		return
	}

	output := string(buf)

	for _, device := range strings.Split(output, "\n") {
		if len(device) > 1 {
			devPath := baseNicPath + device
			destLink, err := filepath.EvalSymlinks(devPath)
			if err == nil {
				if strings.Contains(destLink, "pci") {
					mac, _ := utils.Cat_file(baseNicPath + device + "/address")
					mac = strings.Split(mac, "\n")[0]
					//fmt.Printf("Device %s is physical NIC, and have MAC is %s\n", device, MAC)
					Nic.Items = append(Nic.Items, NicType{IfName: device, MAC: mac})
				}
			}
		}

	}
	return
	/*
		dest, err := os.Readlink("/sys/class/net/lo")
		if err != nil {
			fmt.Println("It's not a symlink!", err)
			return

		} else {
			fmt.Println(dest)
		}
	*/
}
