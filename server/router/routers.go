package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirhmvfx/react-golang-blog/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(app *fiber.App, mongoClient *mongo.Client) {
	postController := controller.NewPostController(mongoClient)
	userController := controller.NewUserController(mongoClient)

	app.Get("/api/posts", postController.GetAllPost)
	app.Get("/api/post/:id", postController.GetPost)
	app.Post("/api/post", postController.CreatePost)
	app.Put("/api/post/:id", postController.UpdatePost)
	app.Delete("/api/post/:id", postController.DeletePost)

	app.Post("/api/register", userController.Register)
	app.Post("/api/login", userController.Login)
}
