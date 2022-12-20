package shortnr

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server interface {
	NewXxh3(c *fiber.Ctx) error
	NewB62(c *fiber.Ctx) error
	GetURLXxh3(c *fiber.Ctx) error
	GetURLB62(c *fiber.Ctx) error
}

type service struct {
	Client *mongo.Client
}

func NewService(client *mongo.Client) Server {
	return &service{Client: client}
}

func (s *service) NewXxh3(c *fiber.Ctx) error {
	return nil
}

func (s *service) NewB62(c *fiber.Ctx) error {
	return nil
}

func (s *service) GetURLXxh3(c *fiber.Ctx) error {
	return nil
}

func (s *service) GetURLB62(c *fiber.Ctx) error {
	return nil
}
