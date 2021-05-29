package helpers

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data)
}


