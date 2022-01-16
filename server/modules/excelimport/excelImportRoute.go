package excelimport

import (
	"datarepository/server/helper/loggermdl"
	"datarepository/server/models"
	"net/http"
	"strings"

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

	searchOption = updateSearchOptionForDateSearch(searchOption)
	contactResponse, err := SearchPersonContactService(searchOption)
	if err != nil {
		loggermdl.LogError(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, contactResponse)
	return
	// loggermdl.LogDebug(paginate)
}

func updateSearchOptionForDateSearch(searchOption models.SearchOption) models.SearchOption {
	if searchOption.SearchBy == "Remark" {
		if strings.HasPrefix(searchOption.SearchText, "dob:") {
			searchOption.SearchText = strings.TrimPrefix(searchOption.SearchText, "dob:")
		}
	}
	return searchOption
}
