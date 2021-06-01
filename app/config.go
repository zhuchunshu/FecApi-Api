package app

import (
	"github.com/hibiken/asynq"
	"github.com/zhuchunshu/FecApi-Api/app/Models/Database"
)

const Version = 1.0

var (
	ApiData Database.PersonalAccessToken
	Jobs *asynq.Client
)
