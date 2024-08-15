package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"setup_go/database"
	"setup_go/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
)

func main() {

	initConfig()

	database.InitDatabase()

	engine := html.New("./views", ".html")
	engine.Reload(true)    // Optional. Default: false
	engine.Debug(false)    // Optional. Default: false
	engine.Layout("embed") // Optional. Default: "embed"
	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		BodyLimit: 300 * 1024 * 1024,
		Views:     engine,
	})
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SetUp GO API " + os.Getenv("ENV") + " " + getVersionNumber())
	})

	port := ":7000"
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "prolocal" {
		port = ":7001"
	}
	router.SetUpRouter(app)

	if err := app.Listen(port); err != nil {
		fmt.Println("error start server ->", err)
	}

}

func initConfig() {

	switch os.Getenv("ENV") {
	case "":
		os.Setenv("ENV", "dev")
		viper.SetConfigName("config")
	default:
		viper.SetConfigName("config")
	}
	fmt.Println("Testbug")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func getVersionNumber() string {
	version := "1.0.0"
	inFile, err := os.Open("./Makefile")
	if err != nil {
		log.Error(err.Error() + `: ` + err.Error())
		return version
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lineVersion := scanner.Text()
		if strings.TrimSpace(lineVersion) != "" {
			listFirstLine := strings.Split(lineVersion, " ")
			version = listFirstLine[len(listFirstLine)-1]
			break
		} else {
			break
		}
	}

	return version
}
