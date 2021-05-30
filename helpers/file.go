package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func WriteFile(path string, str string) bool {
	var f *os.File
	var err error
	f, _ = os.OpenFile(path, os.O_APPEND, 0666)

	//使用完毕，需要关闭文件
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	if err != nil {
		fmt.Println("err = ", err)
		return false
	}
	_, err = f.WriteString(str)
	if err != nil {
		fmt.Println("err = ", err)
	}
	return true
}