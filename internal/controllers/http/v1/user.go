package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavelkorshunov/go-skeleton/internal/modules"
	"go.mongodb.org/mongo-driver/bson"
)

func UserController(c *gin.Context) {
	connection := modules.GetConnection()

	collection := connection.Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal("users not found")
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal("cursor failed")
	}

	c.JSON(http.StatusOK, results)
}
