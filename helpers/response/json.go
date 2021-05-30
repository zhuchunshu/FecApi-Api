package response

import (
	"github.com/zhuchunshu/FecApi-Api/app"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/FecApi-Api/helpers/types"
	"runtime"
	"time"
)
var ServerData string = config.GetServerConfig()
func JsonApi(code int, success bool, data types.Map) types.Map {
	data = types.Map{
		"code":   code,
		"success": success,
		"data":   data,
		"server": types.Map{
			"ServerName": helpers.Json_decode(ServerData, "servername"),
			"version":    app.Version,
			"CpuCore":    runtime.NumCPU(),
			"date":       time.Now(),
			"author":     helpers.Json_decode(ServerData, "author"),
		},
	}

	return data
}

func StringApi(code int, success bool, data string) types.Map {
	return types.Map{
		"code":   code,
		"success": success,
		"data":   data,
		"server": types.Map{
			"ServerName": helpers.Json_decode(ServerData, "servername"),
			"version":    app.Version,
			"CpuCore":    runtime.NumCPU(),
			"date":       time.Now(),
			"author":     helpers.Json_decode(ServerData, "author"),
		},
	}
}
