package shortnr

import (
	"context"
	"errors"

	"github.com/norbux/url-shortnr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveRecord() error {
	return errors.New("not implemented")
}

func GetURL() (string, error) {
	return "", errors.New("not implemented")
}

func NextSeq(client *mongo.Client, database string) (int, error) {
	coll := client.Database(database).Collection("seq")
	filter := bson.D{{"_id", "counter"}}

	var counter models.Seq
	update := bson.D{{"$inc", bson.D{{"seq", 1}}}}
	err := coll.FindOneAndUpdate(context.Background(), filter, update, nil).Decode(&counter)
	if err != nil {
		return 0, err
	}

	// err := coll.FindOne(context.Background(), filter).Decode(&counter)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return 0, errors.New("no conter sequence found in the database")
	// 	}

	// 	return 0, errors.New("failed to get next sequence index")
	// }

	// update := bson.D{{"$inc", bson.D{{"seq", 1}}}}
	// filter = bson.D{{"_id", counter.ID}}
	// result, err := coll.UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	return 0, err
	// }

	return counter.Seq, nil
}
