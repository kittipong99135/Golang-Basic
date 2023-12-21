package routes

import (
	controll "golang_workshop/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	authMiddleware := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "2301223",
		},
	})

	// Route Users
	v1.Post("/user", authMiddleware, controll.CrateUsers)
	v1.Get("/user", controll.ReadUser)
	v1.Get("/user:id", authMiddleware, controll.ReadUsers)
	v1.Put("/user:id", authMiddleware, controll.UpdateUsers)
	v1.Delete("/user:id", authMiddleware, controll.RemoveUsers)
	v1.Get("/look", authMiddleware, controll.SearchUser)
	v1.Get("/genertion", authMiddleware, controll.GenerationUser)
}
