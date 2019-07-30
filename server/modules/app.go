package modules

import (
	"datarepository/server/modules/excelimport"

	"github.com/gin-gonic/gin"
)

// Init - Init
func Init(openRoute *gin.RouterGroup) {
	excelimport.Init(openRoute)
}
