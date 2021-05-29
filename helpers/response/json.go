package response

import (
	"github.com/zhuchunshu/FecApi-Api/app"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/FecApi-Api/helpers/types"
	"runtime"
	"time"
)

func JsonApi(code int,result string,data types.Map) types.Map{
	ServerData:=config.GetServerConfig()
	data = types.Map{
		"code":   code,
		"result": result,
		"data":   data,
		"server": types.Map{
			"ServerName": helpers.Json_decode(ServerData, "servername"),
			"version":    app.Version,
			"CpuCore":    runtime.NumCPU(),
			"date" : time.Now(),
			"author":helpers.Json_decode(ServerData, "author"),
		},
	}

	return data
}