package excelimport

import (
	"datarepository/server/models"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tealeg/xlsx"
)

// ImportContactsService -
func ImportContactsService(excelFileHeader *multipart.FileHeader) error {
	importedExcelPath, err := saveFileInDir(excelFileHeader)
	if err != nil {
		return err
	}
	personContactList, err := readPresonContactsFromFile(importedExcelPath)
	if err != nil {
		return err
	}
	return SaveContactDetailsData(personContactList...)
	// save data to database
}

func readPresonContactsFromFile(importedFilePath string) ([]models.PersonContactDetails, error) {
	personContactList := make([]models.PersonContactDetails, 0)
	xlsFile, err := xlsx.OpenFile(importedFilePath)
	if err != nil {
		return personContactList, err
	}
	if len(xlsFile.Sheets) == 0 && len(xlsFile.Sheets[0].Rows) > 2 {
		return personContactList, errors.New("no data to import")
	}
	// loggermdl.LogDebug("len ", len(xlsFile.Sheets[0].Rows))
	for _, row := range xlsFile.Sheets[0].Rows[1:] {
		person := &models.PersonContactDetails{}
		row.ReadStruct(person)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		personContactList = append(personContactList, *person)
		// loggermdl.LogDebug(row)
	}
	return personContactList, nil
}
func saveFileInDir(excelFileHeader *multipart.FileHeader) (string, error) {
	timestamp := time.Now().Unix()

	destFilePath := filepath.Join(models.Config.DataRepo, strconv.FormatInt(timestamp, 10)+excelFileHeader.Filename)
	err := os.MkdirAll(models.Config.DataRepo, os.ModePerm)
	if err != nil {
		return destFilePath, err
	}

	// loggermdl.LogDebug("dest path", destFilePath)
	destFileWriter, err := os.OpenFile(destFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return destFilePath, err
	}
	sourceFileReader, err := excelFileHeader.Open()
	if err != nil {
		return destFilePath, err
	}

	_, err = io.Copy(destFileWriter, sourceFileReader)
	if err != nil {
		return destFilePath, err
	}
	return destFilePath, nil
}

// SearchPersonContactService -
func SearchPersonContactService(searchOption models.SearchOption) (models.ContactResponse, error) {
	query := bson.D{}
	switch searchOption.SearchBy {
	case "Name":
		query = bson.D{{"name", primitive.Regex{Pattern: searchOption.SearchText, Options: "i"}}}
	case "Address":
		query = bson.D{{"address", primitive.Regex{Pattern: searchOption.SearchText, Options: "i"}}}
	case "Mobile":
		query = bson.D{{"mobile", primitive.Regex{Pattern: searchOption.SearchText, Options: "i"}}}
	case "Tags":
		query = bson.D{{"tags", primitive.Regex{Pattern: searchOption.SearchText, Options: "i"}}}
	case "Remark":
		query = bson.D{{"dob", primitive.Regex{Pattern: searchOption.SearchText, Options: "i"}}}
	}

	paginate := searchOption.Paginate
	return SearchContactDetailsDAO(query, paginate)
}
