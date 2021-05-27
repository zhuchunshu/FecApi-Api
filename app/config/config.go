package config

import "github.com/zhuchunshu/FecApi-Api/helpers"

func GetServerConfig() string {
	data := helpers.ReadFile("config/server.json")
	return data
}
