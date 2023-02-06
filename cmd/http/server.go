package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"os"
	"slurp/internal/core/usecases"
	"slurp/internal/handlers"
	"slurp/internal/repositories"
)

func main() {
	apiConfigUc := usecases.CreateApiConfigurationUseCase{
		ApiRepository: repositories.GcpStorageApiConfigurationRepository{ApiConfigurationBucket: os.Getenv("CONFIGURATION_BUCKET")},
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(favicon.New())

	app.Get("/:apiName", func(c *fiber.Ctx) error {
		// Variable is now immutable
		apiName := utils.CopyString(c.Params("apiName"))
		apiConfiguration, err := apiConfigUc.CreateApiConfiguration(apiName)
		if err != nil {
			log.Printf("An error has occured during API configuration parsing: %v", err)
			return c.Status(500).SendString("Error in API retrieval configurations")
		}
		ctx := usecases.CreateContextUseCase{}.CreateContext(*apiConfiguration)

		usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}.SlurpAPI(ctx)
		return c.SendString(fmt.Sprintf("{\"message\": \"API %s slurped\"}", apiName))
	})
	log.Fatal(app.Listen(":3000"))
}
