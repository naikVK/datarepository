package excelimport

import (
	"context"
	"datarepository/server/dbhelper"
	"datarepository/server/models"

	"datarepository/server/helper/loggermdl"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

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
func SearchContactDetailsDAOold(selector bson.D, paginate models.Paginate) ([]models.PersonContactDetails, error) {
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

	}
	if err = cur.Err(); err != nil {
		loggermdl.LogError(err)
		return personList, err
	}
	cur.Close(dbhelper.DefaultCtx)
	return personList, err
}

// SearchContactDetailsDAO - SearchContactDetailsDAO
func SearchContactDetailsDAO(selector bson.D, paginate models.Paginate) (models.ContactResponse, error) {
	contactResponse := models.ContactResponse{}

	// db.getCollection('person').aggregate([
	// 	{"$match":{ "name": { "$regex": /vivek/, $options: 'i' }}},
	// 	{"$facet":
	// 		{
	// 		"totalContacts":[
	// 			{ "$count": "mycount" }],
	// 		"PersonContactDetails":[
	// 			{"$skip":0},
	// 		{"$limit":2}
	// 		  ]
	// 		}
	// 	},

	// 	{ $unwind: "$totalContacts" }
	// 	{"$project":{"TotalContacts":"$totalContacts", "PersonContactDetails":"$PersonContactDetails" }}
	// 	])
	// Skip:  &paginate.Start,
	// Limit: &paginate.Size,
	matchStage := bson.D{{"$match", selector}}
	facetStage := bson.D{
		{"$facet",
			bson.D{
				{"totalContacts", bson.A{bson.D{{"$count", "mycount"}}}},
				{"PersonContactDetails", bson.A{bson.D{{"$skip", paginate.Start}}, bson.D{{"$limit", paginate.Size}}}},
			},
		},
	}

	unwindStage := bson.D{
		{"$unwind", "$totalContacts"},
	}
	projectStage := bson.D{
		{"$project",
			bson.D{
				{"TotalContacts", "$totalContacts.mycount"},
				{"PersonContactDetails", "$PersonContactDetails"},
			},
		},
	}
	opts := options.Aggregate()
	pipeline := mongo.Pipeline{matchStage, facetStage, unwindStage, projectStage}
	cur, err := dbhelper.DbInstance.Collection("person").Aggregate(context.TODO(), pipeline, opts)
	// cur, err := dbhelper.DbInstance.Collection("person").Find(dbhelper.DefaultCtx, selector, &options.FindOptions{
	// Skip:  &paginate.Start,
	// Limit: &paginate.Size,
	// })
	if err != nil {
		loggermdl.LogError(err)
		return contactResponse, err
	}
	// var data interface{}
	// loggermdl.LogDebug(cur.All(dbhelper.DefaultCtx, cu))
	defer cur.Close(dbhelper.DefaultCtx)
	for cur.Next(dbhelper.DefaultCtx) {
		err = cur.Decode(&contactResponse)
		if err == nil {
			return contactResponse, nil
		}

	}
	if err = cur.Err(); err != nil {
		loggermdl.LogError(err)
		return contactResponse, err
	}
	return contactResponse, err
}
