package helpers

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/zhuchunshu/FecApi-Api/helpers/types"
)

func Json_decode(json string, path string) string {
	value := gjson.Get(json, path)
	return value.String()
}

func JsonToMap(str string) types.Map{
	var tempMap types.Map
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap

}