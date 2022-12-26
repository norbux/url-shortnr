package shortnr

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/norbux/url-shortnr/models"
)

type Server interface {
	GetRecord(hash string) (string, error)
	Short(c *fiber.Ctx) error
	GetURL(c *fiber.Ctx) error
}

type service struct {
	Client   *mongo.Client
	Database string
}

func NewService(client *mongo.Client, db string) Server {
	return &service{Client: client, Database: db}
}

func (s *service) Short(c *fiber.Ctx) error {
	req := new(models.NewHashRequest)
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	if req.Method < 1 {
		return fiber.ErrBadRequest
	}

	nextSeq, err := s.NextSeq()
	if err != nil {
		return err
	}

	var hash string

	switch req.Method {
	case 1:
		hash, err = Xxh3Hash(strconv.Itoa(nextSeq))
		if err != nil {
			return err
		}
	case 2:
		hash, err = B62Hash(nextSeq)
		if err != nil {
			return err
		}
	}

	record := models.URLMap{LongURL: req.URL, Hash: hash}
	err = s.SaveRecord(&record)
	if err != nil {
		return err
	}

	resp := models.NewHashResponse{URL: req.URL, Hash: hash}

	return c.JSON(resp)
}

func (s *service) GetURL(c *fiber.Ctx) error {
	hash := c.Params("hash")
	matched, err := regexp.MatchString(`\b[0-9a-fA-F]+\b`, hash)
	if err != nil {
		return errors.New("bad hash input")
	}

	if len(hash) < 1 || !matched {
		return errors.New("empty or invalid hash string")
	}

	url, err := s.GetRecord(hash)
	if err != nil {
		return err
	}

	return c.Redirect(url, fiber.StatusFound)
}
