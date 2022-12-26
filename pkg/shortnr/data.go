package shortnr

import (
	"context"
	"errors"
	"log"

	"github.com/norbux/url-shortnr/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) SaveRecord(urlMap *models.URLMap) error {
	if len(urlMap.LongURL) < 1 || len(urlMap.Hash) < 1 {
		return errors.New("the URL value is empty")
	}

	coll := s.Client.Database(s.Database).Collection("url_map")
	result, err := coll.InsertOne(context.Background(), urlMap)
	if err != nil {
		return err
	}

	log.Printf("%v", result.InsertedID)

	return nil
}

func (s *service) GetRecord(hash string) (string, error) {
	if len(hash) < 1 {
		return "", errors.New("hash can't be empty string")
	}

	resp := new(models.URLMap)
	coll := s.Client.Database(s.Database).Collection("url_map")
	filter := bson.D{{"hash", hash}}
	err := coll.FindOne(context.Background(), filter, nil).Decode(&resp)
	if err != nil {
		return "", err
	}

	return resp.LongURL, nil
}

func (s *service) NextSeq() (int, error) {
	coll := s.Client.Database(s.Database).Collection("seq")
	filter := bson.D{{"_id", "counter"}}

	var counter models.Seq
	update := bson.D{{"$inc", bson.D{{"seq", 1}}}}
	err := coll.FindOneAndUpdate(context.Background(), filter, update, nil).Decode(&counter)
	if err != nil {
		return 0, err
	}

	return counter.Seq, nil
}
