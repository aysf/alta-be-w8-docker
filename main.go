package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitDB() {
	godotenv.Load()

	connectionString2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString2), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func InitMigration() {
	DB.AutoMigrate(&User{})
}

func init() {
	InitDB()
	InitMigration()
}

func main() {

	e := echo.New()

	e.GET("/", hello)
	e.POST("/user", createUser)
	e.GET("/users", getUsers)

	e.Start(":8080")

}

func hello(c echo.Context) error {
	godotenv.Load()

	name := os.Getenv("NAME")

	greeting := fmt.Sprintf("Hello %v", name)

	return c.String(http.StatusOK, greeting)
}

func createUser(c echo.Context) error {
	newUser := new(User)
	c.Bind(newUser)

	if err := DB.Create(newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    newUser,
	})
}

func getUsers(c echo.Context) error {
	users := new([]User)

	if err := DB.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}
