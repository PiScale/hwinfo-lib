package utils

import (
	"os/exec"
)

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
