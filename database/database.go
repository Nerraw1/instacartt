package database

import (
	"fmt"
	"instacartt/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var DNS = "host=localhost user=postgres password=admin dbname=Users port=5432 sslmode=disable"

func Migration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("not connected to the database")
	}
	fmt.Print("connecteed to the database")
	DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Images{})
}

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stars       string `json:"stars"`
	Quantity    int    `json:"quantity"`
	Location    string `json:"location"`
}

//images
type Images struct {
	Image_id   uint   `json:"image_id" gorm:"primaryKey"`
	Image_name string ` json:"imag_name"`
	Image_url  string `json:"image_url"`
}

//foodlistapi
func AddProduct(c *fiber.Ctx) error {
	var food Product
	if err := c.BodyParser(&food); err != nil {
		return err

	}
	DB.Create(&food)
	return c.JSON(&food)
}

//images
func AddImages(c *fiber.Ctx) error {
	var je Images
	if err := c.BodyParser(&je); err != nil {
		return err

	}
	DB.Create(&je)
	return c.JSON(&je)
}
