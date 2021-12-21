package main

import (
	"datarepository/server/helper"
	"datarepository/server/helper/loggermdl"
	"datarepository/server/models"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"datarepository/server/dbhelper"
	"datarepository/server/modules"
	"os"
	"path/filepath"

	"go.uber.org/zap/zapcore"

	"github.com/gin-gonic/gin"
)

var serverDirPath string
var configDirPath string

func main() {
	ginEngine := gin.New()

	err := InitAll(ginEngine)
	if err != nil {
		loggermdl.LogError(err)
		return
	}
	err = ginEngine.Run(":" + models.Config.AppPort)
	if err != nil {
		loggermdl.LogError(err)
		return
	}
}

// InitAll -
func InitAll(ginEngine *gin.Engine) error {
	loggermdl.Init("logs/datarepository.log", 0, 1, 0, zapcore.DebugLevel)
	_, err := helper.InitConfig("config/config.toml", &models.Config)
	if err != nil {
		loggermdl.LogError(err)
		return err
	}

	serverDirPath, _ := os.Getwd()
	cfgDirPath := filepath.Join(serverDirPath, "config")
	configDirPath = cfgDirPath
	dbhelper.Init(filepath.Join(configDirPath, "mongo-config.toml"), "dataRepoHost")
	ginEngine.Use(CORSMiddleware())
	openRouteGroup := ginEngine.Group("o")
	ginEngine.StaticFS("/", http.Dir("../dist"))
	modules.Init(openRouteGroup)
	loggermdl.LogDebug("appport", models.Config.AppPort)
	openbrowser("http://localhost:" + models.Config.AppPort)
	return nil
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		println("darwin")
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		loggermdl.LogError(err)
	}
}

// CORSMiddleware -
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
