package excelimport

import (
	"datarepository/server/dbhelper"
	"datarepository/server/models"

	"go.mongodb.org/mongo-driver/bson"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveContactDetailsData - SaveContactDetailsData
func SaveContactDetailsData(personList ...models.PersonContactDetails) error {
	val := make([]interface{}, 0)
	for _, data := range personList {
		// data.MixField = data.Name + " " + data.Mobile + " " + data.Tag + " " + data.Address
		val = append(val, data)
	}
	// val = append(val, personList)

	_, err := dbhelper.DbInstance.Collection("person").InsertMany(dbhelper.DefaultCtx, val, &options.InsertManyOptions{})
	return err
}

// SearchContactDetailsDAO - SearchContactDetailsDAO
func SearchContactDetailsDAO(selector bson.D, paginate models.Paginate) ([]models.PersonContactDetails, error) {
	personList := []models.PersonContactDetails{}
	loggermdl.LogDebug(selector)

	// Skip:  &paginate.Start,
	// Limit: &paginate.Size,
	cur, err := dbhelper.DbInstance.Collection("person").Find(dbhelper.DefaultCtx, selector, &options.FindOptions{
		// Skip:  &paginate.Start,
		// Limit: &paginate.Size,
	})
	if err != nil {
		loggermdl.LogError(err)
		return personList, err
	}
	// var data interface{}
	// loggermdl.LogDebug(cur.All(dbhelper.DefaultCtx, cu))
	for cur.Next(dbhelper.DefaultCtx) {

		var personDetails models.PersonContactDetails
		err = cur.Decode(&personDetails)
		if err == nil {
			personList = append(personList, personDetails)
		}
		// loggermdl.LogDebug(personDetails)
	}
	if err = cur.Err(); err != nil {
		loggermdl.LogError(err)
		return personList, err
	}
	cur.Close(dbhelper.DefaultCtx)
	return personList, err
}
