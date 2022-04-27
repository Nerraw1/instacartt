package main

import (
	"errors"
	"fmt"
	"instacartt/database"
	"instacartt/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}

func Routes(app *fiber.App) {

	//users
	app.Post("/user", route.Registration)
	app.Get("/user", route.GetUsers)
	//foodlist
	app.Post("/product", AddProduct)
	app.Get("/product", GetProductName)
	app.Get("/product/:id", GetProduct)
	app.Delete("/product/:id", Delete)
	app.Put("/product/:id", Update)
	//image
	app.Get("/images", GetImages)
	app.Get("/images/:id", GetImage)

}

//foodlist
func AddProduct(c *fiber.Ctx) error {
	var food database.Product
	if err := c.BodyParser(&food); err != nil {
		return c.SendString(err.Error())
	}

	database.DB.Create(&food)
	return c.JSON(&food)
}

func GetProductName(c *fiber.Ctx) error {
	var food []database.Product

	database.DB.Find(&food)
	return c.JSON(&food)
}

func FindProduct(id int, food *database.Product) error {
	database.DB.Find(&food, "id=?", id)
	if food.ID == 0 {
		return errors.New("ProductId does not existed")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var food database.Product
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(&food)

}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var food database.Product
	if err != nil {
		return c.SendString(err.Error())
	}
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}
	database.DB.Delete(&food)
	return c.SendString("Deleted product")
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var food database.Product
	if err != nil {
		return c.SendString(err.Error())
	}
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}
	if err := c.BodyParser(&food); err != nil {
		return c.SendString(err.Error())
	}
	database.DB.Save(&food)
	return c.JSON(&food)
}

//images
func GetImages(c *fiber.Ctx) error {
	je := []database.Images{}
	database.DB.Find(&je)
	return c.JSON(je)
}

func GetImage(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	je := database.Images{}
	database.DB.Find(&je, "image_id", id)
	return c.JSON(je)

}

func main() {
	database.Migration()
	app := fiber.New()
	Routes(app)
	log.Fatal(app.Listen(":3000"))
	fmt.Println("InstaCart")
	app.Get("/", Welcome)
}
