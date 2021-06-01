/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package V1

import (
	"net/http"
	"net/url"
)

const (
	TypeEmailDelivery   = "email:deliver"
)

func QQHook(urls string,qq string,title string,token string) {
	_, _ = http.PostForm(urls+token+"/private/"+qq+"/"+title, url.Values{})
}