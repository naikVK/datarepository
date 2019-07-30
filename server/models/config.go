package models

// AppConfig -
type AppConfig struct {
	AppName                     string
	AppVersion                  string
	AppPort                     string
	DatabaseDir                 string
	LogDir                      string
	MaxSessionPerUser           int
	SessionExpireTimeInMinutes  int
	SchedularTimeIntervalInSecs int
	NoofBackups                 int
	DataRepo                    string
}

var (
	// Config -
	Config AppConfig
)
