package excelimport

import (
	"datarepository/server/models"
	"net/http"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
	"github.com/gin-gonic/gin"
)

// Init - Init
func Init(openGroup *gin.RouterGroup) {
	openGroup.POST("importContacts", ImportContactsRoute)
	openGroup.POST("searchContact", SearchContactRoute)

}

// ImportContactsRoute - ImportContactsRoute
func ImportContactsRoute(c *gin.Context) {
	loggermdl.LogDebug("In: importContacts")
	file, err := c.FormFile("file")
	if nil != err {
		loggermdl.LogError("error uploading file: ", err)
		c.JSON(http.StatusBadRequest, "error uploading file")
		return
	}
	// excelFile, err := file.Open()

	if nil != err {
		loggermdl.LogError("error uploading file: ", err)
		c.JSON(http.StatusBadRequest, "error uploading file")
		return
	}
	err = ImportContactsService(file)
	if nil != err {
		loggermdl.LogError("error uploading data: ", err)
		c.JSON(http.StatusBadRequest, "error uploading data")
	}
	c.JSON(http.StatusOK, "SUCCESS")
	return
}

// SearchContactRoute - SearchContactRoute
func SearchContactRoute(c *gin.Context) {
	searchOption := models.SearchOption{}
	// paginate := models.Paginate{}
	err := c.BindJSON(&searchOption)
	// err := c.BindJSON(&paginate)

	if err != nil {
		loggermdl.LogError(err)
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	personList, err := SearchPersonContactService(searchOption)
	if err != nil {
		loggermdl.LogError(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, personList)
	return
	// loggermdl.LogDebug(paginate)
}
