/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package String

import (
	"encoding/base64"
	"log"
)



func Base64Encode(text []byte)string  {
	return base64.StdEncoding.EncodeToString(text)
}

func Base64DeCode(text string)[]byte{
	decode,err :=base64.StdEncoding.DecodeString(text)
	if err!=nil {
		log.Fatalln(err)
	}
	return decode
}