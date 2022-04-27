package route

import (
	"instacartt/database"
	"instacartt/model"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c *fiber.Ctx) error {
	var user model.User
	new_user := new(model.User)

	if err := c.BodyParser(&new_user); err != nil {
		return c.Status(500).SendString("Server error")
	}
	regEmail := regexp.MustCompile("[a=zA-Z0-9_]+@[yahoogmail]+[.][com]{3}")
	formatterEmail := regEmail.MatchString(new_user.Email)
	database.DB.Find(&user, "email=?", new_user.Email)
	database.DB.Find(&user, "username=?", new_user.Username)
	uniqueEmail := new_user.Email != user.Email
	uniqueUsername := new_user.Username != user.Username
	usernameLength := len(new_user.Username) >= 8
	passwordLength := len(new_user.Password) >= 8
	hash, _ := HashPassword(new_user.Password)
	new_user.Password = hash

	if formatterEmail && uniqueEmail && uniqueUsername && usernameLength && passwordLength {
		database.DB.Create(&new_user)
	} else {
		if !formatterEmail {
			return c.SendString("Email not valid")
		}
		if !uniqueEmail {
			return c.SendString("Email already exist!")
		}
		if !uniqueUsername {
			return c.SendString("Username already exist!")
		}
		if !usernameLength {
			return c.SendString("Username length should be atleast 8 characters!")
		}
		if !passwordLength {
			return c.SendString("Password length should be atleast 8 characters!")
		}
	}

	return c.JSON(&fiber.Map{
		"message": "User successfully registered",
		"USER":    new_user,
	})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GetUsers(c *fiber.Ctx) error {
	var user []model.User
	database.DB.Find(&user)
	if len(user) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User Does not Exist!",
		})
	}
	return c.JSON(&fiber.Map{
		"Users": user,
	})
}
