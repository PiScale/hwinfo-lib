package utils

import (
	"io/ioutil"
	//"os/exec"
)

func Cat_file(filepath string) (ret string, err error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	ret = string(buf)
	return
}

/*
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
