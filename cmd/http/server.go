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
	"slurp/internal/handlers/repositories"
)

func main() {
	createContextUc := usecases.CreateContextUseCase{
		ApiConfigurationRepository: repositories.GcpStorageApiConfigurationRepository{ApiConfigurationBucket: os.Getenv("CONFIGURATION_BUCKET")},
	}
	slurpAnApiUc := usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(favicon.New())

	app.Get("/:apiName", func(c *fiber.Ctx) error {
		// Variable is now immutable
		apiName := utils.CopyString(c.Params("apiName"))
		ctx, err := createContextUc.CreateContext(apiName)
		if err != nil {
			log.Printf("An error has occured during API configuration parsing: %v", err)
			return c.Status(500).SendString("Error in API retrieval configurations")
		}

		slurpAnApiUc.SlurpAPI(*ctx)
		return c.SendString(fmt.Sprintf("{\"message\": \"API %s slurped\"}", apiName))
	})
	log.Fatal(app.Listen(":3000"))
}
