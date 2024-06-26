package controller

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sirhmvfx/react-golang-blog/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostController struct {
	collection *mongo.Collection
}

func NewPostController(client *mongo.Client) *PostController {
	collection := client.Database("react-golang-blog").Collection("posts")
	return &PostController{collection}
}

func ExtractedUserId(c *fiber.Ctx) (primitive.ObjectID, error) {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID, err := primitive.ObjectIDFromHex(claims["userID"].(string))
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userID, nil
}

func (p *PostController) GetAllPost(c *fiber.Ctx) error {
	var posts []model.Post

	cursor, err := p.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var post model.Post
		if err := cursor.Decode(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to decode post"})
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cursor error"})
	}

	return c.JSON(posts)
}

func (p *PostController) CreatePost(c *fiber.Ctx) error {
	var post model.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse body request"})
	}

	userID, err := ExtractedUserId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to extract user id"})
	}

	post.Author = userID

	if post.Title == "" || post.Body == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Title and body field is required"})
	}

	post.ID = primitive.NewObjectID()
	_, err = p.collection.InsertOne(context.Background(), post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

func (p *PostController) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Post"})
	}

	var post model.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}

	_, err = p.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": post})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update post"})
	}

	return c.JSON(post)
}

func (p *PostController) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post Id"})
	}

	var post model.Post
	err = p.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "post not found"})
	}

	return c.JSON(post)
}

func (p *PostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}

	_, err = p.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete Post"})
	}

	return c.JSON(fiber.Map{"success": "Successfully deleted post"})
}
