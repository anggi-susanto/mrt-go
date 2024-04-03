package main

import (
	"context"

	"github.com/gofiber/swagger"

	"github.com/anggi-susanto/mrt-go/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"

	_ "github.com/anggi-susanto/mrt-go/docs"
	mongoRepo "github.com/anggi-susanto/mrt-go/internal/repository/mongo"
	"github.com/anggi-susanto/mrt-go/internal/rest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title MRT Waste Water API
// @version 1.0
// @description This is an API Document for MRT Waste Water
// @contact.name MRT WW Support
// @contact.email antscpk06@gmail.com
// @license.name GPL-3.0 License
// @license.url https://www.gnu.org/licenses/gpl-3.0.html
// @host localhost:3000
// @BasePath /
func main() {
	config := config.Config{
		MongoConfig: config.MongoConfig{
			Uri:                  "mongodb://localhost:27017",
			Database:             "mrt",
			WasteWaterCollection: "waste_water",
		},
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoConfig.Uri))
	if err != nil {
		logrus.Fatal(err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			logrus.Fatal(err)
		}
	}()

	// Start the server
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("MRT API is UP and RUNNING!")
	})

	wasteWaterRepo := mongoRepo.NewWasteWaterRepository(mongoClient, &config.MongoConfig)
	rest.NewWasteWaterHandler(app, wasteWaterRepo)

	logrus.Fatal(app.Listen(":3000"))

}
