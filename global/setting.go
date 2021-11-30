package global

import (
	"Pro/blog-service/pkg/logger"
	"Pro/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
