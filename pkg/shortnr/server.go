package shortnr

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/norbux/url-shortnr/models"
)

type Server interface {
	NewXxh3(c *fiber.Ctx) error
	NewB62(c *fiber.Ctx) error
	GetURLXxh3(c *fiber.Ctx) error
	GetURLB62(c *fiber.Ctx) error
}

type service struct {
	Client   *mongo.Client
	Database string
}

func NewService(client *mongo.Client, db string) Server {
	return &service{Client: client, Database: db}
}

func (s *service) NewXxh3(c *fiber.Ctx) error {
	nextSeq, err := NextSeq(s.Client, s.Database)
	if err != nil {
		return err
	}

	hash, err := Xxh3Hash(strconv.Itoa(nextSeq))
	if err != nil {
		return err
	}

	resp := models.HashResponse{Hash: hash}

	return c.JSON(resp)
}

func (s *service) NewB62(c *fiber.Ctx) error {
	nextSeq, err := NextSeq(s.Client, s.Database)
	if err != nil {
		return err
	}

	hash, err := B62Hash(nextSeq)
	if err != nil {
		return err
	}

	resp := models.HashResponse{Hash: hash}

	return c.JSON(resp)
}

func (s *service) GetURLXxh3(c *fiber.Ctx) error {
	return nil
}

func (s *service) GetURLB62(c *fiber.Ctx) error {
	return nil
}
