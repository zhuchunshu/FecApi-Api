package helpers

import "github.com/tidwall/gjson"

func Json_decode(json string, path string) string {
	value := gjson.Get(json, path)
	return value.String()
}
